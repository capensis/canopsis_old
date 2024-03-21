package techmetrics

//go:generate mockgen -destination=../../../mocks/lib/techmetrics/sender.go git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/techmetrics Sender

import (
	"context"
	"math"
	"sync"
	"time"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

const maxUrlLength = 500

type Sender interface {
	Run(ctx context.Context)

	SendFifoEvent(metric FifoEventMetric)

	SendCpsEvent(metric CpsEventMetric)

	SendQueue(metricName string, timestamp time.Time, length int64)

	SendCheEntityInfo(timestamp time.Time, name string)
	SendCheEvent(metric CheEventMetric)

	SendAxePeriodical(metric AxePeriodicalMetric)
	SendAxeEvent(metric AxeEventMetric)

	SendPBehaviorPeriodical(metric PbehaviorPeriodicalMetric)

	SendApiRequest(metric ApiRequestMetric)

	SendCorrelationEvent(metric CorrelationEventMetric)
	SendCorrelationRetries(metric CorrelationRetriesMetric)

	SendDynamicInfosEvent(metric DynamicInfoEventMetric)

	SendActionEvent(metric ActionEventMetric)
}

func NewSender(
	configProvider config.TechMetricsConfigProvider,
	interval time.Duration,
	poolRetryCount int,
	poolRetryTimeout time.Duration,
	logger zerolog.Logger,
) Sender {
	return &sender{
		configProvider:   configProvider,
		interval:         interval,
		poolRetryCount:   poolRetryCount,
		poolRetryTimeout: poolRetryTimeout,
		logger:           logger,

		batches: make(map[string][][]any),
	}
}

type sender struct {
	configProvider config.TechMetricsConfigProvider
	interval       time.Duration
	logger         zerolog.Logger

	poolRetryCount   int
	poolRetryTimeout time.Duration

	batchesMx sync.Mutex
	batches   map[string][][]any

	pool postgres.Pool
}

func (s *sender) Run(ctx context.Context) {
	defer func() {
		if s.pool != nil {
			s.pool.Close()
		}
	}()

	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.send(ctx)
		}
	}
}

func (s *sender) SendQueue(metricName string, timestamp time.Time, length int64) {
	s.addBatch(metricName, []any{
		timestamp.UTC(),
		length,
	})
}

func (s *sender) SendAxePeriodical(metric AxePeriodicalMetric) {
	s.addBatch(AxePeriodical, []any{
		metric.Timestamp.UTC(),
		metric.Interval.Microseconds(),
		metric.Events,
		metric.IdleEvents,
	})
}

func (s *sender) SendPBehaviorPeriodical(metric PbehaviorPeriodicalMetric) {
	s.addBatch(PBehaviorPeriodical, []any{
		metric.Timestamp.UTC(),
		metric.Interval.Microseconds(),
		metric.Events,
		metric.Entities,
		metric.Pbehaviors,
	})
}

func (s *sender) SendCheEntityInfo(timestamp time.Time, name string) {
	s.addBatch(CheInfos, []any{
		timestamp.UTC(),
		name,
	})
}

func (s *sender) SendApiRequest(metric ApiRequestMetric) {
	url := metric.Url
	if len(url) > maxUrlLength {
		url = url[:maxUrlLength]
	}
	s.addBatch(ApiRequests, []any{
		metric.Timestamp.UTC(),
		metric.Method,
		url,
		metric.Interval.Microseconds(),
	})
}

func (s *sender) SendCpsEvent(metric CpsEventMetric) {
	s.addBatch(CanopsisEvent, []any{
		metric.Timestamp.UTC(),
		metric.EventType,
		metric.Interval.Microseconds(),
		metric.IsOkState,
	})
}

func (s *sender) SendFifoEvent(metric FifoEventMetric) {
	var externalRequests map[string]int64
	if len(metric.ExternalRequests) > 0 {
		externalRequests = metric.ExternalRequests
	}

	s.addBatch(FIFOEvent, []any{
		metric.Timestamp.UTC(),
		metric.EventType,
		metric.Interval.Microseconds(),
		externalRequests,
	})
}

func (s *sender) SendCheEvent(metric CheEventMetric) {
	var externalRequests map[string]int64
	if len(metric.ExternalRequests) > 0 {
		externalRequests = metric.ExternalRequests
	}

	s.addBatch(CheEvent, []any{
		metric.Timestamp.UTC(),
		metric.EventType,
		metric.Interval.Microseconds(),
		metric.EntityType,
		metric.IsNewEntity,
		metric.IsInfosUpdated,
		metric.IsServicesUpdated,
		metric.IsStateSettingUpdated,
		metric.ExecutedEnrichRules,
		externalRequests,
	})
}

func (s *sender) SendAxeEvent(metric AxeEventMetric) {
	s.addBatch(AxeEvent, []any{
		metric.Timestamp.UTC(),
		metric.EventType,
		metric.Interval.Microseconds(),
		metric.EntityType,
		metric.AlarmChangeType,
		metric.IsOkState,
		metric.IsCountersUpdated,
	})
}

func (s *sender) SendCorrelationEvent(metric CorrelationEventMetric) {
	s.addBatch(CorrelationEvent, []any{
		metric.Timestamp.UTC(),
		metric.EventType,
		metric.Interval.Microseconds(),
		metric.MatchedRuleCount,
		metric.MatchedRuleTypes,
	})
}

func (s *sender) SendCorrelationRetries(metric CorrelationRetriesMetric) {
	s.addBatch(CorrelationRetries, []any{
		metric.Timestamp.UTC(),
		metric.Type,
		metric.Retries,
	})
}

func (s *sender) SendDynamicInfosEvent(metric DynamicInfoEventMetric) {
	s.addBatch(DynamicInfosEvent, []any{
		metric.Timestamp.UTC(),
		metric.EventType,
		metric.Interval.Microseconds(),
		metric.ExecutedRules,
	})
}

func (s *sender) SendActionEvent(metric ActionEventMetric) {
	s.addBatch(ActionEvent, []any{
		metric.Timestamp.UTC(),
		metric.EventType,
		metric.Interval.Microseconds(),
		metric.ExecutedRules,
		metric.ExecutedWebhooks,
	})
}

func (s *sender) send(ctx context.Context) {
	if !s.configProvider.Get().Enabled {
		if s.pool != nil {
			s.pool.Close()
			s.pool = nil
		}

		s.cleanBatches()
		return
	}

	batches := s.flushBatches()

	if s.pool == nil {
		var err error
		s.pool, err = postgres.NewTechMetricsPool(ctx, s.poolRetryCount, s.poolRetryTimeout)
		if err != nil {
			s.logger.Err(err).Msg("cannot connect tech metrics Postgres")
			return
		}
	}

	bulkSize := canopsis.DefaultBulkSize
	for metricName, rows := range batches {
		columns := s.getColumns(metricName)
		if len(columns) == 0 {
			s.logger.Error().Msgf("unknown columns for %q", metricName)
			continue
		}

		rowsCount := len(rows)
		bulkCount := int(math.Ceil(float64(rowsCount) / float64(bulkSize)))
		for i := 0; i < bulkCount; i++ {
			begin := i * bulkSize
			end := (i + 1) * bulkSize
			if end > rowsCount {
				end = rowsCount
			}
			_, err := s.pool.CopyFrom(ctx, pgx.Identifier{metricName}, columns, pgx.CopyFromRows(rows[begin:end]))
			if err != nil {
				s.logger.Err(err).Msg("cannot send tech metrics")
				return
			}
		}
	}
}

func (s *sender) flushBatches() map[string][][]any {
	s.batchesMx.Lock()
	defer s.batchesMx.Unlock()

	batches := make(map[string][][]any, len(s.batches))
	for metricName, items := range s.batches {
		if len(items) == 0 {
			continue
		}
		batches[metricName] = make([][]any, 0, len(items))
	}
	res := s.batches
	s.batches = batches

	return res
}

func (s *sender) cleanBatches() {
	s.batchesMx.Lock()
	defer s.batchesMx.Unlock()

	clear(s.batches)
}

func (s *sender) addBatch(metricName string, args []any) {
	if !s.configProvider.Get().Enabled {
		return
	}

	s.batchesMx.Lock()
	defer s.batchesMx.Unlock()

	if _, ok := s.batches[metricName]; !ok {
		s.batches[metricName] = make([][]any, 0, 1)
	}

	s.batches[metricName] = append(s.batches[metricName], args)
}

func (s *sender) getColumns(metricName string) []string {
	switch metricName {
	case FIFOEvent:
		return []string{
			"time",
			"type",
			"interval",
			"external_requests",
		}
	case CheEvent:
		return []string{
			"time",
			"type",
			"interval",
			"entity_type",
			"is_new_entity",
			"is_infos_updated",
			"is_services_updated",
			"is_state_settings_updated",
			"executed_enrich_rules",
			"external_requests",
		}
	case AxeEvent:
		return []string{
			"time",
			"type",
			"interval",
			"entity_type",
			"alarm_change_type",
			"is_ok_state",
			"is_counters_updated",
		}
	case CanopsisEvent:
		return []string{
			"time",
			"type",
			"interval",
			"is_ok_state",
		}
	case CorrelationEvent:
		return []string{
			"time",
			"type",
			"interval",
			"matched_rules",
			"matched_rule_types",
		}
	case ActionEvent:
		return []string{
			"time",
			"type",
			"interval",
			"executed_rules",
			"executed_webhooks",
		}
	case DynamicInfosEvent:
		return []string{
			"time",
			"type",
			"interval",
			"executed_rules",
		}
	case AxePeriodical:
		return []string{
			"time",
			"interval",
			"events",
			"idle_events",
		}
	case PBehaviorPeriodical:
		return []string{
			"time",
			"interval",
			"events",
			"entities",
			"pbehaviors",
		}
	case FIFOQueue:
		return []string{
			"time",
			"length",
		}
	case CheInfos:
		return []string{
			"time",
			"name",
		}
	case ApiRequests:
		return []string{
			"time",
			"method",
			"url",
			"interval",
		}
	case CorrelationRetries:
		return []string{
			"time",
			"type",
			"retries",
		}
	}

	return nil
}
