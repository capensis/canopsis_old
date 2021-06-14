package canopsis

import (
	"git.canopsis.net/canopsis/go-engines/lib/mongo"
	"time"
)

// Globals
const (
	EnvCpsAMQPTestExchange          = "canopsis.tests"
	ActionEngineName                = "engine-action"
	ActionQueueName                 = "Engine_action"
	ActionAxeRPCClientQueueName     = "Engine_action_axe_rpc_client"
	ActionWebhookRPCClientQueueName = "Engine_action_webhook_rpc_client"
	ActionConsumerName              = "action"
	ActionRPCConsumerName           = "action_rpc"
	AxeExchangeName                 = "engine-axe"
	AxeQueueName                    = "Engine_axe"
	AxeServiceRPCClientQueueName    = "Engine_axe_service_rpc_client"
	AxePbehaviorRPCClientQueueName  = "Engine_axe_pbehavior_rpc_client"
	AxeRPCQueueServerName           = "Engine_axe_rpc_server"
	AxeConsumerName                 = "axe"
	AxeRPCConsumerName              = "axe_rpc"
	CheExchangeName                 = "canopsis.events"
	CheEngineName                   = "engine-che"
	CheQueueName                    = "Engine_che"
	CheConsumerName                 = "che"
	DbName                          = mongo.DB
	DefaultBulkSize                 = 1000
	DefaultEventAuthor              = "system"
	DoneAutosolveDelay              = 15 * 60
	DynamicInfosQueueName           = "Engine_dynamic_infos"
	DynamicInfosConsumerName        = "dynamic-infos"
	HeartBeatExchangeName           = "canopsis.events"
	HeartBeatQueueName              = "Engine_heartbeat"
	HeartBeatConsumerName           = "heartbeat"
	MaxPythonTimestamp              = 253402297199 // compat: max timestamp in python
	StatsExchangeName               = "canopsis.events"
	StatsQueueName                  = "Engine_stat"
	StatsConsumerName               = "stat"
	StatsDatabase                   = "canopsis"
	StatsMeasurement                = "event_state_history"
	StatsngExchangeName             = "amq.direct"
	StatsngQueueName                = "Engine_statsng"
	PBehaviorEngineName             = "engine-pbehavior"
	PBehaviorQueueName              = "Engine_pbehavior"
	PBehaviorRPCQueueServerName     = "Engine_pbehavior_rpc_server"
	PBehaviorConsumerName           = "pbehavior"
	PBehaviorRPCConsumerName        = "pbehavior_rpc"
	PluginExtension                 = ".so"
	ServiceEngineName               = "engine-service"
	ServiceQueueName                = "Engine_service"
	ServiceRPCQueueServerName       = "Engine_service_rpc_server"
	ServiceConsumerName             = "service"
	ServiceRPCConsumerName          = "service_rpc"
	WebhookRPCQueueServerName       = "Engine_webhook_rpc_server"
	WebhookRPCConsumerName          = "webhook_rpc"
	FIFOExchangeName                = ""
	FIFOQueueName                   = "Engine_fifo"
	FIFOAckExchangeName             = ""
	FIFOAckQueueName                = "FIFO_ack"
	FIFOConsumerName                = "fifo"
	CorrelationQueueName            = "Engine_correlation"
	CorrelationConsumerName         = "correlation"
	PeriodicalWaitTime              = time.Minute
	JsonContentType                 = "application/json"
	CanopsisEventsExchange          = "canopsis.events"
)
