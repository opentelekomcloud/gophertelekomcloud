package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts is a struct which represents the parameters of update function
type UpdateOpts struct {
	Topics []UpdateItem `json:"topics" required:"true"`
}

// UpdateItem represents the object of one topic in update function
type UpdateItem struct {
	// Name can not be updated
	Name string `json:"id" required:"true"`
	// Aging time in hour.
	RetentionTime *int `json:"retention_time,omitempty"`
	// Whether synchronous replication is enabled.
	SyncReplication *bool `json:"sync_replication,omitempty"`
	// Whether synchronous flushing is enabled.
	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`
	// Number of the partitions.
	NewPartitionNumbers *int `json:"new_partition_numbers,omitempty"`
	// Specifying brokers for new partitions.
	NewPartitionBrokers []int `json:"new_partition_brokers,omitempty"`
	// Topic configuration.
	TopicOtherConfigs []TopicOtherConfigs `json:"topic_other_configs,omitempty"`
	// Topic description.
	Description string `json:"topic_desc,omitempty"`
}

// Update is a method which can be able to update topics
func Update(client *golangsdk.ServiceClient, instanceID string, opts UpdateOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	_, err = client.Put(client.ServiceURL("instances", instanceID, "topics"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
