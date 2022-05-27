/*
Package executor contains operation executors.
*/
package executor

import (
	"context"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/metrics"
	operationlib "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/operation"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/utils"
)

// NewAckExecutor creates new executor.
func NewAckExecutor(metricsSender metrics.Sender, configProvider config.AlarmConfigProvider) operationlib.Executor {
	return &ackExecutor{
		metricsSender:  metricsSender,
		configProvider: configProvider,
	}
}

type ackExecutor struct {
	metricsSender  metrics.Sender
	configProvider config.AlarmConfigProvider
}

// Exec creates new ack step for alarm.
func (e *ackExecutor) Exec(
	ctx context.Context,
	operation types.Operation,
	alarm *types.Alarm,
	_ *types.Entity,
	time types.CpsTime,
	userID, role, initiator string,
) (types.AlarmChangeType, error) {
	params := operation.Parameters

	if userID == "" {
		userID = params.User
	}

	doubleAck := false
	if alarm.Value.ACK != nil {
		doubleAck = true
	}

	if doubleAck && !e.configProvider.Get().AllowDoubleAck {
		return "", nil
	}

	err := alarm.PartialUpdateAck(
		time,
		params.Author,
		utils.TruncateString(params.Output, e.configProvider.Get().OutputLength),
		userID,
		role,
		initiator,
	)

	if err != nil {
		return "", err
	}

	if !doubleAck {
		go func() {
			metricsUserID := ""
			if initiator == types.InitiatorUser {
				metricsUserID = userID
			}
			e.metricsSender.SendAck(context.Background(), *alarm, metricsUserID, time.Time)
		}()

		return types.AlarmChangeTypeAck, nil
	}

	return types.AlarmChangeTypeDoubleAck, nil
}
