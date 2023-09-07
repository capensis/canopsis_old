package axe

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	libamqp "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/amqp"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/alarm"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/encoding"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/engine"
	libentity "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/entity"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/entityservice/statecounters"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/metrics"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/operation"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/pbehavior"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/rpc"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/mongo"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type rpcPBehaviorClientMessageProcessor struct {
	DbClient                 mongo.DbClient
	MetricsSender            metrics.Sender
	PublishCh                libamqp.Channel
	RemediationRpc           engine.RPCClient
	Executor                 operation.Executor
	EntityAdapter            libentity.Adapter
	AlarmAdapter             alarm.Adapter
	PbehaviorAdapter         pbehavior.Adapter
	StateCountersService     statecounters.StateCountersService
	Decoder                  encoding.Decoder
	Encoder                  encoding.Encoder
	Logger                   zerolog.Logger
	FeaturePrintEventOnError bool
}

func (p *rpcPBehaviorClientMessageProcessor) Process(ctx context.Context, msg engine.RPCMessage) error {
	data := strings.Split(msg.CorrelationID, "**")
	if len(data) != 2 {
		return fmt.Errorf("RPC PBehavior Client: bad correlation_id: %s", msg.CorrelationID)
	}

	correlationId := data[0]
	routingKey := data[1]

	var replyEvent []byte
	var event rpc.PbehaviorResultEvent
	err := p.Decoder.Decode(msg.Body, &event)
	if err != nil || event.Alarm == nil || event.Entity == nil {
		p.logError(err, "RPC PBehavior Client: invalid event", msg.Body)

		return p.publishResult(ctx, routingKey, correlationId, p.getErrRpcEvent(fmt.Errorf("invalid event")))
	}

	alarmChangeType := types.AlarmChangeTypeNone

	if event.PbhEvent.EventType != "" {
		var updatedServiceInfos map[string]statecounters.UpdatedServicesInfo
		firstTimeTran := true
		var alarmChange types.AlarmChange

		err = p.DbClient.WithTransaction(ctx, func(tCtx context.Context) error {
			if !firstTimeTran {
				var alarmsWithEntity []types.AlarmWithEntity
				err := p.AlarmAdapter.GetOpenedAlarmsWithEntityByAlarmIDs(ctx, []string{event.Alarm.ID}, &alarmsWithEntity)
				if err != nil {
					return err
				}

				if len(alarmsWithEntity) == 0 {
					return fmt.Errorf("entity with id = %s or alarm with id = %s is not found after transaction rollback", event.Entity.ID, event.Alarm.ID)
				}

				event.Entity = &alarmsWithEntity[0].Entity
				event.Alarm = &alarmsWithEntity[0].Alarm
			}

			firstTimeTran = false

			alarmChange = types.NewAlarmChangeByAlarm(*event.Alarm)

			alarmChange.Type, err = p.Executor.Exec(
				ctx,
				types.Operation{
					Type: event.PbhEvent.EventType,
					Parameters: types.OperationParameters{
						PbehaviorInfo: &event.PbhEvent.PbehaviorInfo,
						Author:        event.PbhEvent.Author,
						Output:        event.PbhEvent.Output,
					},
				},
				event.Alarm,
				event.Entity,
				event.PbhEvent.Timestamp,
				"",
				"",
				types.InitiatorSystem,
			)

			if err != nil {
				return err
			}

			p.updateEntity(ctx, event.PbhEvent.Entity, *event.Alarm, alarmChange.Type)
			p.updatePbhLastAlarmDate(ctx, alarmChange.Type, event.Alarm.Value.PbehaviorInfo)

			updatedServiceInfos, err = p.StateCountersService.UpdateServiceCounters(tCtx, *event.Entity, event.Alarm, alarmChange)
			return err
		})
		if err != nil {
			if engine.IsConnectionError(err) {
				return err
			}

			p.logError(err, "RPC PBehavior Client: cannot update alarm", msg.Body)
			return p.publishResult(ctx, routingKey, correlationId, p.getErrRpcEvent(fmt.Errorf("cannot update alarm: %v", err)))
		}

		// services alarms
		go func() {
			for servID, servInfo := range updatedServiceInfos {
				err := p.StateCountersService.UpdateServiceState(context.Background(), servID, servInfo)
				if err != nil {
					p.Logger.Err(err).Msg("failed to update service state")
				}
			}
		}()

		p.MetricsSender.SendEventMetrics(
			*event.Alarm,
			*event.Entity,
			alarmChange,
			event.PbhEvent.Timestamp.Time,
			types.InitiatorSystem,
			"",
			"",
		)

		if p.RemediationRpc != nil {
			body, err := p.Encoder.Encode(types.RPCRemediationEvent{
				Alarm:       event.Alarm,
				Entity:      event.PbhEvent.Entity,
				AlarmChange: alarmChange,
			})
			if err != nil {
				p.logError(err, "RPC PBehavior Client: failed to encode rpc call to engine-remediation", msg.Body)
			} else {
				err = p.RemediationRpc.Call(ctx, engine.RPCMessage{
					CorrelationID: utils.NewID(),
					Body:          body,
				})
				if err != nil {
					if engine.IsConnectionError(err) {
						return err
					}

					p.logError(err, "RPC PBehavior Client: failed to send rpc call to engine-remediation", msg.Body)
				}
			}
		}
	}

	replyEvent, err = p.getRpcEvent(rpc.AxeResultEvent{
		Alarm:           event.Alarm,
		AlarmChangeType: alarmChangeType,
		Error:           nil,
	})
	if err != nil {
		p.logError(err, "RPC PBehavior Client: failed to create rpc result event", msg.Body)

		replyEvent = p.getErrRpcEvent(errors.New("failed to create rpc result event"))
	}

	err = p.publishResult(ctx, routingKey, correlationId, replyEvent)
	if err != nil {
		p.logError(err, "RPC PBehavior Client: cannot sent message result back to sender", msg.Body)

		return err
	}

	return nil
}

func (p *rpcPBehaviorClientMessageProcessor) publishResult(ctx context.Context, routingKey string, correlationID string, event []byte) error {
	return p.PublishCh.PublishWithContext(
		ctx,
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: correlationID,
			Body:          event,
			DeliveryMode:  amqp.Persistent,
		},
	)
}

func (p *rpcPBehaviorClientMessageProcessor) logError(err error, errMsg string, msg []byte) {
	if p.FeaturePrintEventOnError {
		p.Logger.Debug().Err(err).Str("event", string(msg)).Msg(errMsg)
	} else {
		p.Logger.Err(err).Msg(errMsg)
	}
}

func (p *rpcPBehaviorClientMessageProcessor) getErrRpcEvent(err error) []byte {
	msg, _ := p.getRpcEvent(rpc.AxeResultEvent{Error: &rpc.Error{Error: err}})
	return msg
}

func (p *rpcPBehaviorClientMessageProcessor) getRpcEvent(event rpc.AxeResultEvent) ([]byte, error) {
	msg, err := p.Encoder.Encode(event)
	if err != nil {
		p.Logger.Err(err).Msg("cannot encode event")
		return nil, err
	}

	return msg, nil
}

func (p *rpcPBehaviorClientMessageProcessor) updateEntity(ctx context.Context, entity *types.Entity, alarm types.Alarm, changeType types.AlarmChangeType) {
	switch changeType {
	case types.AlarmChangeTypeCreateAndPbhEnter, types.AlarmChangeTypePbhEnter,
		types.AlarmChangeTypePbhLeave, types.AlarmChangeTypePbhLeaveAndEnter:
		entity.PbehaviorInfo = alarm.Value.PbehaviorInfo
		err := p.EntityAdapter.UpdatePbehaviorInfo(ctx, entity.ID, entity.PbehaviorInfo)
		if err != nil {
			p.Logger.Err(err).Msg("cannot update entity")
		}
	}
}

func (p *rpcPBehaviorClientMessageProcessor) updatePbhLastAlarmDate(ctx context.Context, changeType types.AlarmChangeType, pbehaviorInfo types.PbehaviorInfo) {
	if changeType != types.AlarmChangeTypeCreateAndPbhEnter &&
		changeType != types.AlarmChangeTypePbhEnter &&
		changeType != types.AlarmChangeTypePbhLeaveAndEnter {
		return
	}

	err := p.PbehaviorAdapter.UpdateLastAlarmDate(ctx, pbehaviorInfo.ID, types.CpsTime{Time: time.Now()})
	if err != nil {
		p.Logger.Err(err).Msg("cannot update pbehavior")
	}
}
