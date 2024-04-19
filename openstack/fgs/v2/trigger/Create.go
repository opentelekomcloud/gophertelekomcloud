package trigger

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	FuncUrn         string     `json:"-"`
	TriggerTypeCode string     `json:"trigger_type_code" required:"true"`
	TriggerStatus   string     `json:"trigger_status,omitempty"`
	EventTypeCode   string     `json:"event_type_code,omitempty"`
	EventData       *EventData `json:"event_data" required:"true"`
}

type EventData struct {
	Name                string           `json:"name,omitempty"`
	ScheduleType        string           `json:"schedule_type,omitempty"`
	Schedule            string           `json:"schedule,omitempty"`
	UserEvent           string           `json:"user_event,omitempty"`
	Type                *int             `json:"type,omitempty"`
	Path                string           `json:"path,omitempty"`
	Protocol            string           `json:"protocol,omitempty"`
	ReqMethod           string           `json:"req_method,omitempty"`
	GroupID             string           `json:"group_id,omitempty"`
	GroupName           string           `json:"group_name,omitempty"`
	MatchMode           string           `json:"match_mode,omitempty"`
	EnvName             string           `json:"env_name,omitempty"`
	EnvID               string           `json:"env_id,omitempty"`
	Auth                string           `json:"auth,omitempty"`
	FuncInfo            *TriggerFuncInfo `json:"func_info,omitempty"`
	SlDomain            string           `json:"sl_domain,omitempty"`
	BackendType         string           `json:"backend_type,omitempty"`
	Operations          []string         `json:"operations,omitempty"`
	InstanceID          string           `json:"instance_id,omitempty"`
	CollectionName      string           `json:"collection_name,omitempty"`
	DbName              string           `json:"db_name,omitempty"`
	DbPassword          string           `json:"db_password,omitempty"`
	BatchSize           *int             `json:"batch_size,omitempty"`
	QueueID             string           `json:"queue_id,omitempty"`
	ConsumerGroupID     string           `json:"consumer_group_id,omitempty"`
	PollingInterval     *int             `json:"polling_interval,omitempty"`
	StreamName          string           `json:"stream_name,omitempty"`
	SharditeratorType   string           `json:"sharditerator_type,omitempty"`
	PollingUnit         string           `json:"polling_unit,omitempty"`
	MaxFetchBytes       *int             `json:"max_fetch_bytes,omitempty"`
	IsSerial            string           `json:"is_serial,omitempty"`
	LogGroupID          string           `json:"log_group_id,omitempty"`
	LogTopicID          string           `json:"log_topic_id,omitempty"`
	Bucket              string           `json:"bucket,omitempty"`
	Prefix              string           `json:"prefix,omitempty"`
	Suffix              string           `json:"suffix,omitempty"`
	Events              []string         `json:"events,omitempty"`
	TopicUrn            string           `json:"topic_urn,omitempty"`
	TopicIds            []string         `json:"topic_ids,omitempty"`
	KafkaUser           string           `json:"kafka_user,omitempty"`
	KafkaPassword       string           `json:"kafka_password,omitempty"`
	KafkaConnectAddress string           `json:"kafka_connect_address,omitempty"`
	KafkaSSLEnable      *bool            `json:"kafka_ssl_enable,omitempty"`
	AccessPassword      string           `json:"access_password,omitempty"`
	AccessUser          string           `json:"access_user,omitempty"`
	ConnectAddress      string           `json:"connect_address,omitempty"`
	ExchangeName        string           `json:"exchange_name,omitempty"`
	Vhost               string           `json:"vhost,omitempty"`
	SSLEnable           *bool            `json:"ssl_enable,omitempty"`
	// return values
	TriggerId     string   `json:"trigger_id,omitempty"`
	InvokeUrl     string   `json:"invoke_url,omitempty"`
	RomaAppId     string   `json:"roma_app_id,omitempty"`
	InstanceAddrs []string `json:",omitempty"`
	Mode          string   `json:"mode,omitempty"`
}

type TriggerFuncInfo struct {
	FunctionUrn    string `json:"function_urn,omitempty"`
	InvocationType string `json:"invocation_type,omitempty"`
	Timeout        int    `json:"timeout" required:"true"`
	Version        string `json:"version,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*TriggerFuncResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("fgs", "triggers", opts.FuncUrn), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res TriggerFuncResp
	return &res, extract.Into(raw.Body, &res)
}

type TriggerFuncResp struct {
	TriggerId       string     `json:"trigger_id"`
	TriggerTypeCode string     `json:"trigger_type_code"`
	TriggerStatus   string     `json:"trigger_status"`
	EventData       *EventData `json:"event_data"`
	LastUpdatedTime string     `json:"last_updated_time"`
	CreatedTime     string     `json:"created_time"`
}
