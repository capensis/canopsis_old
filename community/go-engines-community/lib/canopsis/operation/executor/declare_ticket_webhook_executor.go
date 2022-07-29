package executor

import (
	"context"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/metrics"
	operationlib "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/operation"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/utils"
)

// NewDeclareTicketWebhookExecutor creates new executor.
func NewDeclareTicketWebhookExecutor(configProvider config.AlarmConfigProvider, metricsSender metrics.Sender) operationlib.Executor {
	return &declareTicketWebhookExecutor{configProvider: configProvider, metricsSender: metricsSender}
}

type declareTicketWebhookExecutor struct {
	configProvider config.AlarmConfigProvider
	metricsSender  metrics.Sender
}

// Exec creates new declare ticket step for alarm.
func (e *declareTicketWebhookExecutor) Exec(
	_ context.Context,
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

	err := alarm.PartialUpdateDeclareTicket(
		time,
		params.Author,
		utils.TruncateString(params.Output, e.configProvider.Get().OutputLength),
		params.Ticket,
		params.Data,
		userID,
		role,
		initiator,
	)
	if err != nil {
		return "", err
	}

	go func() {
		metricsUserID := ""
		if initiator == types.InitiatorUser {
			metricsUserID = userID
		}
		e.metricsSender.SendTicket(context.Background(), *alarm, metricsUserID, time.Time)
	}()

	return types.AlarmChangeTypeDeclareTicketWebhook, nil
}
