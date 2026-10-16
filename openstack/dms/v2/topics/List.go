package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts contains the query parameters of the List function.
type ListOpts struct {
	// Offset, which is the position where the query starts. The value must be greater than or equal to 0.
	Offset string `q:"offset,omitempty"`
	// Maximum number of instances returned in the current query. The default value is 10. The value ranges from 1 to 50.
	Limit string `q:"limit,omitempty"`
}

// List is used to query details about topics of a Kafka instance.
// Send GET /v2/{project_id}/instances/{instance_id}/topics
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*ListResponse, error) {
	if instanceID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "instanceID"}
	}

	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "topics").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// ListResponse is a struct that contains the list response
type ListResponse struct {
	// Total number of topics.
	Total int `json:"total"`
	// Maximum number of records to be displayed on a page.
	Size int `json:"size"`
	// Number of remaining partitions.
	RemainPartitions int `json:"remain_partitions"`
	// Total number of partitions.
	MaxPartitions int `json:"max_partitions"`
	// Maximum number of partitions in a single topic.
	TopicMaxPartitions int `json:"topic_max_partitions"`
	// Topic list.
	Topics []Topic `json:"topics"`
}

type Topic struct {
	// Whether this policy is the default policy.
	PoliciesOnly bool `json:"policiesOnly"`
	// Topic name.
	Name string `json:"name"`
	// Number of replicas, which is configured to ensure data reliability.
	Replication int `json:"replication"`
	// Number of topic partitions, which is used to set the number of concurrently consumed messages.
	Partition int `json:"partition"`
	// Retention period of a message.
	RetentionTime int `json:"retention_time"`
	// Whether synchronous replication is enabled. After this function is enabled,
	// the acks parameter on the producer client must be set to -1.
	// Otherwise, this parameter does not take effect.
	SyncReplication bool `json:"sync_replication"`
	// Whether synchronous flushing is enabled. The default value is false.
	// Synchronous flushing compromises performance.
	SyncMessageFlush bool `json:"sync_message_flush"`
	// Extended configuration.
	ExternalConfigs map[string]interface{} `json:"external_configs"`
	// Topic type. Options:
	//    0: common topic
	//    1: system (internal) topic
	TopicType int `json:"topic_type"`
	// Other topic configurations.
	TopicOtherConfigs []TopicOtherConfigsResp `json:"topic_other_configs"`
	// Topic description.
	Description string `json:"topic_desc"`
	// Topic creation time.
	CreatedAt int64 `json:"created_at"`
}

type TopicOtherConfigsResp struct {
	// Configuration name.
	Name string `json:"name"`
	// Valid value.
	ValidValues string `json:"valid_values"`
	// Default value.
	DefaultValue string `json:"default_value"`
	// Configuration type. The value can be dynamic or static.
	ConfigType string `json:"config_type"`
	// Configuration value.
	Value string `json:"value"`
	// Configuration value type.
	ValueType string `json:"value_type"`
}
