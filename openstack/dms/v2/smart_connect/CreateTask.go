package smart_connect

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateTaskOpts struct {
	// Smart Connect task name
	TaskName string `json:"task_name,omitempty"`
	// Indicates whether to start a task later
	StartLater *bool `json:"start_later,omitempty"`
	// Topic of a Smart Connect task
	Topics string `json:"topics,omitempty"`
	// Regular expression of the topic of a Smart Connect task
	TopicsRegex string `json:"topics_regex,omitempty"`
	// Source type of a Smart Connect task
	SourceType string `json:"source_type,omitempty"`
	// Source configuration of a Smart Connect task
	SourceTask *SmartConnectTaskSourceConfig `json:"source_task,omitempty"`
	// Target type of a Smart Connect task
	SinkType string `json:"sink_type,omitempty"`
	// Target configuration of a Smart Connect task
	SinkTask *SmartConnectTaskSinkConfig `json:"sink_task,omitempty"`
}

type SmartConnectTaskSourceConfig struct {
	// Redis-specific fields
	RedisAddress     string `json:"redis_address,omitempty"`
	RedisType        string `json:"redis_type,omitempty"`
	DcsInstanceId    string `json:"dcs_instance_id,omitempty"`
	RedisPassword    string `json:"redis_password,omitempty"`
	SyncMode         string `json:"sync_mode,omitempty"`
	FullSyncWaitMs   int    `json:"full_sync_wait_ms,omitempty"`
	FullSyncMaxRetry int    `json:"full_sync_max_retry,omitempty"`
	Ratelimit        int    `json:"ratelimit,omitempty"`
	// Kafka-specific fields
	CurrentClusterName         string `json:"current_cluster_name,omitempty"`
	ClusterName                string `json:"cluster_name,omitempty"`
	UserName                   string `json:"user_name,omitempty"`
	Password                   string `json:"password,omitempty"`
	SaslMechanism              string `json:"sasl_mechanism,omitempty"`
	InstanceId                 string `json:"instance_id,omitempty"`
	BootstrapServers           string `json:"bootstrap_servers,omitempty"`
	SecurityProtocol           string `json:"security_protocol,omitempty"`
	Direction                  string `json:"direction,omitempty"`
	SyncConsumerOffsetsEnabled *bool  `json:"sync_consumer_offsets_enabled,omitempty"`
	ReplicationFactor          int    `json:"replication_factor,omitempty"`
	TaskNum                    int    `json:"task_num,omitempty"`
	RenameTopicEnabled         *bool  `json:"rename_topic_enabled,omitempty"`
	ProvenanceHeaderEnabled    *bool  `json:"provenance_header_enabled,omitempty"`
	ConsumerStrategy           string `json:"consumer_strategy,omitempty"`
	CompressionType            string `json:"compression_type,omitempty"`
	TopicsMapping              string `json:"topics_mapping,omitempty"`
}

type SmartConnectTaskSinkConfig struct {
	// Redis-specific fields
	RedisAddress  string `json:"redis_address,omitempty"`
	RedisType     string `json:"redis_type,omitempty"`
	DcsInstanceId string `json:"dcs_instance_id,omitempty"`
	RedisPassword string `json:"redis_password,omitempty"`

	// OBS-specific fields
	ConsumerStrategy    string `json:"consumer_strategy,omitempty"`
	DestinationFileType string `json:"destination_file_type,omitempty"`
	DeliverTimeInterval int    `json:"deliver_time_interval,omitempty"`
	AccessKey           string `json:"access_key,omitempty"`
	SecretKey           string `json:"secret_key,omitempty"`
	ObsBucketName       string `json:"obs_bucket_name,omitempty"`
	ObsPath             string `json:"obs_path,omitempty"`
	PartitionFormat     string `json:"partition_format,omitempty"`
	RecordDelimiter     string `json:"record_delimiter,omitempty"`
	StoreKeys           *bool  `json:"store_keys,omitempty"`
}

// This API is used to create a Smart Connect task.
// POST /v2/{project_id}/instances/{instance_id}/connector/tasks
func CreateTask(client *golangsdk.ServiceClient, instanceId string, opts CreateTaskOpts) (*CreateTaskResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", instanceId, "connector", "tasks"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateTaskResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateTaskResponse struct {
	TaskName    string                        `json:"task_name"`
	Topics      string                        `json:"topics"`
	TopicsRegex string                        `json:"topics_regex"`
	SourceType  string                        `json:"source_type"`
	SourceTask  *SmartConnectTaskSourceConfig `json:"source_task"`
	SinkType    string                        `json:"sink_type"`
	SinkTask    *SmartConnectTaskSinkConfig   `json:"sink_task"`
	ID          string                        `json:"id"`
	Status      string                        `json:"status"`
	CreateTime  int64                         `json:"create_time"`
}
