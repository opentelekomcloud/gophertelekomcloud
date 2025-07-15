package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List all topics belong to the instance id
func List(client *golangsdk.ServiceClient, instanceID string) (*ListResponse, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceID, "topics"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// ListResponse is a struct that contains the list response
type ListResponse struct {
	Total              int     `json:"total"`
	Size               int     `json:"size"`
	RemainPartitions   int     `json:"remain_partitions"`
	MaxPartitions      int     `json:"max_partitions"`
	TopicMaxPartitions int     `json:"topic_max_partitions"`
	Topics             []Topic `json:"topics"`
}

type Topic struct {
	PoliciesOnly      bool                    `json:"policiesOnly"`
	Name              string                  `json:"name"`
	Replication       int                     `json:"replication"`
	Partition         int                     `json:"partition"`
	RetentionTime     int                     `json:"retention_time"`
	SyncReplication   bool                    `json:"sync_replication"`
	SyncMessageFlush  bool                    `json:"sync_message_flush"`
	ExternalConfigs   interface{}             `json:"external_configs"`
	TopicType         int                     `json:"topic_type"`
	TopicOtherConfigs []TopicOtherConfigsResp `json:"topic_other_configs"`
	Description       string                  `json:"topic_desc"`
	CreatedAt         int64                   `json:"created_at"`
}

type TopicOtherConfigsResp struct {
	Name         string `json:"name"`
	ValidValues  string `json:"valid_values"`
	DefaultValue string `json:"default_value"`
	ConfigType   string `json:"config_type"`
	Value        string `json:"value"`
	ValueType    string `json:"value_type"`
}
