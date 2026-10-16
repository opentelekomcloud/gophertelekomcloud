package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts is a struct that contains all the parameters of create function
type CreateOpts struct {
	// Topic name. A topic name consists of 3 to 200 characters, starts with a letter,
	// and contains only letters, hyphens (-), underscores (_), periods (.), and digits.
	Name string `json:"id" required:"true"`
	// Number of replicas, which is configured to ensure data reliability.
	// Value range: 1 to 3.
	Replication int `json:"replication,omitempty"`
	// Whether synchronous flushing is enabled. The default value is false.
	// Synchronous flushing compromises performance.
	SyncMessageFlush bool `json:"sync_message_flush,omitempty"`
	// Number of topic partitions, which is used to set the number of concurrently
	// consumed messages. Value range: 1-200.
	Partition int `json:"partition,omitempty"`
	// Whether synchronous replication is enabled. After this function is enabled,
	// the acks parameter on the producer client must be set to -1.
	// Otherwise, this parameter does not take effect.
	SyncReplication bool `json:"sync_replication,omitempty"`
	// Indicates the retention period of a message. The default value is 72.
	// The value range is 1-720, in hours.
	RetentionTime int `json:"retention_time,omitempty"`
	// Topic configuration.
	TopicOtherConfigs []TopicOtherConfigs `json:"topic_other_configs,omitempty"`
	// Topic description.
	Description string `json:"topic_desc,omitempty"`
}

type TopicOtherConfigs struct {
	// Configuration name.
	Name string `json:"name,omitempty"`
	// Configuration value.
	Value string `json:"value,omitempty"`
}

// Create a kafka topic with given parameters
// Send POST /v2/{project_id}/instances/{instance_id}/topics
func Create(client *golangsdk.ServiceClient, instanceID string, opts CreateOpts) (*CreateResponse, error) {
	if instanceID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "instanceID"}
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", instanceID, "topics"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// CreateResponse is a struct that contains the create response
type CreateResponse struct {
	// Topic name.
	Name string `json:"name"`
}
