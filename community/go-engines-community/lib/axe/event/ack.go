package event

import (
	"context"
	"errors"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/entitycounters"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/entitycounters/calculator"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/metrics"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/rpc"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/mongo"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	mongodriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewAckProcessor(
	client mongo.DbClient,
	configProvider config.AlarmConfigProvider,
	entityServiceCountersCalculator calculator.EntityServiceCountersCalculator,
	eventsSender entitycounters.EventsSender,
	metaAlarmPostProcessor MetaAlarmPostProcessor,
	metricsSender metrics.Sender,
	logger zerolog.Logger,
) Processor {
	return &ackProcessor{
		client:                          client,
		alarmCollection:                 client.Collection(mongo.AlarmMongoCollection),
		entityCollection:                client.Collection(mongo.EntityMongoCollection),
		configProvider:                  configProvider,
		entityServiceCountersCalculator: entityServiceCountersCalculator,
		eventsSender:                    eventsSender,
		metaAlarmPostProcessor:          metaAlarmPostProcessor,
		metricsSender:                   metricsSender,
		logger:                          logger,
	}
}

type ackProcessor struct {
	client                          mongo.DbClient
	alarmCollection                 mongo.DbCollection
	entityCollection                mongo.DbCollection
	configProvider                  config.AlarmConfigProvider
	entityServiceCountersCalculator calculator.EntityServiceCountersCalculator
	eventsSender                    entitycounters.EventsSender
	metaAlarmPostProcessor          MetaAlarmPostProcessor
	metricsSender                   metrics.Sender
	logger                          zerolog.Logger
}

func (p *ackProcessor) Process(ctx context.Context, event rpc.AxeEvent) (Result, error) {
	result := Result{}
	if event.Entity == nil || !event.Entity.Enabled {
		return result, nil
	}

	entity := event.Entity
	match := getOpenAlarmMatchWithStepsLimit(event)
	match["v.ack"] = nil
	conf := p.configProvider.Get()
	newStepQuery := stepUpdateQueryWithInPbhInterval(types.AlarmStepAck, event.Parameters.Output, event.Parameters)
	update := []bson.M{
		{"$set": bson.M{
			"v.ack":              newStepQuery,
			"v.steps":            addStepUpdateQuery(newStepQuery),
			"v.last_update_date": event.Parameters.Timestamp,
		}},
		{"$unset": bson.A{
			"not_acked_metric_type",
			"not_acked_metric_send_time",
			"not_acked_since",
		}},
	}
	var updatedServiceStates map[string]entitycounters.UpdatedServicesInfo
	notAckedMetricType := ""

	err := p.client.WithTransaction(ctx, func(ctx context.Context) error {
		result = Result{}
		updatedServiceStates = nil
		notAckedMetricType = ""

		doubleAck := false
		beforeAlarm := types.Alarm{}
		opts := options.FindOneAndUpdate().
			SetReturnDocument(options.Before).
			SetProjection(bson.M{
				"not_acked_metric_type":      1,
				"not_acked_metric_send_time": 1,
			})
		err := p.alarmCollection.FindOneAndUpdate(ctx, match, update, opts).Decode(&beforeAlarm)
		if err != nil {
			if !errors.Is(err, mongodriver.ErrNoDocuments) {
				return err
			}

			if !conf.AllowDoubleAck {
				return nil
			}

			doubleAck = true
			delete(match, "v.ack")
			err = p.alarmCollection.FindOneAndUpdate(ctx, match, update, opts).Decode(&beforeAlarm)
			if err != nil {
				if errors.Is(err, mongodriver.ErrNoDocuments) {
					return nil
				}

				return err
			}
		}

		if beforeAlarm.NotAckedMetricSendTime != nil {
			notAckedMetricType = beforeAlarm.NotAckedMetricType
		}

		alarm := types.Alarm{}
		err = p.alarmCollection.FindOne(ctx, bson.M{"_id": beforeAlarm.ID}).Decode(&alarm)
		if err != nil {
			if errors.Is(err, mongodriver.ErrNoDocuments) {
				return nil
			}
			return err
		}

		alarmChange := types.NewAlarmChange()
		if doubleAck {
			alarmChange.Type = types.AlarmChangeTypeDoubleAck
		} else {
			alarmChange.Type = types.AlarmChangeTypeAck
		}

		result.Forward = true
		result.Alarm = alarm
		result.AlarmChange = alarmChange

		if event.Parameters.IdleRuleApply != "" {
			result.Entity, err = updateEntityByID(ctx, entity.ID, bson.M{"$set": bson.M{
				"last_idle_rule_apply": event.Parameters.IdleRuleApply,
			}}, p.entityCollection)
			if err != nil {
				return err
			}
		}

		result.IsCountersUpdated, updatedServiceStates, err = p.entityServiceCountersCalculator.CalculateCounters(ctx, entity, &result.Alarm, result.AlarmChange)

		return err
	})

	if err != nil || result.Alarm.ID == "" {
		return result, err
	}

	go p.postProcess(context.Background(), event, result, updatedServiceStates, notAckedMetricType)

	return result, nil
}

func (p *ackProcessor) postProcess(
	ctx context.Context,
	event rpc.AxeEvent,
	result Result,
	updatedServiceStates map[string]entitycounters.UpdatedServicesInfo,
	notAckedMetricType string,
) {
	p.metricsSender.SendEventMetrics(
		result.Alarm,
		*event.Entity,
		result.AlarmChange,
		event.Parameters.Timestamp.Time,
		event.Parameters.Initiator,
		event.Parameters.User,
		event.Parameters.Instruction,
		notAckedMetricType,
	)

	for servID, servInfo := range updatedServiceStates {
		err := p.eventsSender.UpdateServiceState(ctx, servID, servInfo)
		if err != nil {
			p.logger.Err(err).Msg("failed to update service state")
		}
	}

	err := p.metaAlarmPostProcessor.Process(ctx, event, rpc.AxeResultEvent{
		Alarm:           &result.Alarm,
		AlarmChangeType: result.AlarmChange.Type,
	})
	if err != nil {
		p.logger.Err(err).Msg("cannot process meta alarm")
	}
}
