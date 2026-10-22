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

// SmartConnectTaskSourceConfig is the source configuration of a create request (SmartConnectTaskReqSourceConfig).
// Task responses use SmartConnectTaskRespSourceConfig.
type SmartConnectTaskSourceConfig struct {
	// Redis source (REDIS_REPLICATOR_SOURCE)
	RedisAddress     string `json:"redis_address,omitempty"`
	RedisType        string `json:"redis_type,omitempty"`
	DcsInstanceId    string `json:"dcs_instance_id,omitempty"`
	RedisPassword    string `json:"redis_password,omitempty"`
	SyncMode         string `json:"sync_mode,omitempty"`
	FullSyncWaitMs   int    `json:"full_sync_wait_ms,omitempty"`
	FullSyncMaxRetry int    `json:"full_sync_max_retry,omitempty"`
	Ratelimit        int    `json:"ratelimit,omitempty"`

	// Kafka source (KAFKA_REPLICATOR_SOURCE)
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

// SmartConnectTaskSinkConfig is the target configuration of a create request (SmartConnectTaskReqSinkConfig).
// Task responses use SmartConnectTaskRespSinkConfig.
type SmartConnectTaskSinkConfig struct {
	// Redis target (REDIS_REPLICATOR_SINK)
	RedisAddress  string `json:"redis_address,omitempty"`
	RedisType     string `json:"redis_type,omitempty"`
	DcsInstanceId string `json:"dcs_instance_id,omitempty"`
	RedisPassword string `json:"redis_password,omitempty"`

	// OBS target (OBS_SINK)
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

// SmartConnectTaskRespSourceConfig is the source configuration returned in task responses.
// Unlike the request it has no redis_password and password.
type SmartConnectTaskRespSourceConfig struct {
	// Redis source (REDIS_REPLICATOR_SOURCE)
	RedisAddress     string `json:"redis_address"`
	RedisType        string `json:"redis_type"`
	DcsInstanceId    string `json:"dcs_instance_id"`
	SyncMode         string `json:"sync_mode"`
	FullSyncWaitMs   int    `json:"full_sync_wait_ms"`
	FullSyncMaxRetry int    `json:"full_sync_max_retry"`
	Ratelimit        int    `json:"ratelimit"`

	// Kafka source (KAFKA_REPLICATOR_SOURCE)
	CurrentClusterName         string `json:"current_cluster_name"`
	ClusterName                string `json:"cluster_name"`
	UserName                   string `json:"user_name"`
	SaslMechanism              string `json:"sasl_mechanism"`
	InstanceId                 string `json:"instance_id"`
	BootstrapServers           string `json:"bootstrap_servers"`
	SecurityProtocol           string `json:"security_protocol"`
	Direction                  string `json:"direction"`
	SyncConsumerOffsetsEnabled bool   `json:"sync_consumer_offsets_enabled"`
	ReplicationFactor          int    `json:"replication_factor"`
	TaskNum                    int    `json:"task_num"`
	RenameTopicEnabled         bool   `json:"rename_topic_enabled"`
	ProvenanceHeaderEnabled    bool   `json:"provenance_header_enabled"`
	ConsumerStrategy           string `json:"consumer_strategy"`
	CompressionType            string `json:"compression_type"`
	TopicsMapping              string `json:"topics_mapping"`
}

// SmartConnectTaskRespSinkConfig is the target configuration returned in task responses.
// Unlike the request it has no redis_password, access_key and secret_key, but adds target_db and OBS defaults.
type SmartConnectTaskRespSinkConfig struct {
	// Redis target (REDIS_REPLICATOR_SINK)
	RedisAddress  string `json:"redis_address"`
	RedisType     string `json:"redis_type"`
	DcsInstanceId string `json:"dcs_instance_id"`
	// Target database, -1 by default. A pointer keeps database 0 distinguishable from an absent field.
	TargetDB *int `json:"target_db"`

	// OBS target (OBS_SINK)
	ConsumerStrategy     string `json:"consumer_strategy"`
	DestinationFileType  string `json:"destination_file_type"`
	DeliverTimeInterval  int    `json:"deliver_time_interval"`
	ObsBucketName        string `json:"obs_bucket_name"`
	ObsPath              string `json:"obs_path"`
	PartitionFormat      string `json:"partition_format"`
	RecordDelimiter      string `json:"record_delimiter"`
	StoreKeys            bool   `json:"store_keys"`
	ObsPartSize          int    `json:"obs_part_size"`
	FlushSize            int    `json:"flush_size"`
	Timezone             string `json:"timezone"`
	SchemaGeneratorClass string `json:"schema_generator_class"`
	PartitionerClass     string `json:"partitioner_class"`
	ValueConverter       string `json:"value_converter"`
	KeyConverter         string `json:"key_converter"`
	KvDelimiter          string `json:"kv_delimiter"`
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
	TaskName    string                            `json:"task_name"`
	Topics      string                            `json:"topics"`
	TopicsRegex string                            `json:"topics_regex"`
	SourceType  string                            `json:"source_type"`
	SourceTask  *SmartConnectTaskRespSourceConfig `json:"source_task"`
	SinkType    string                            `json:"sink_type"`
	SinkTask    *SmartConnectTaskRespSinkConfig   `json:"sink_task"`
	ID          string                            `json:"id"`
	Status      string                            `json:"status"`
	CreateTime  int64                             `json:"create_time"`
}
