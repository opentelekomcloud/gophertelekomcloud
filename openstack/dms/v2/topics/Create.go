package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOps is a struct that contains all the parameters of create function
type CreateOpts struct {
	// the name/ID of a topic
	Name string `json:"id" required:"true"`
	// topic replications, value range: 1-3
	Replication int `json:"replication,omitempty"`
	// Whether synchronous flushing is enabled. The default value is false.
	// Synchronous flushing compromises performance.
	SyncMessageFlush bool `json:"sync_message_flush,omitempty"`
	// topic partitions, value range: 1-100
	Partition int `json:"partition,omitempty"`
	// Number of topic partitions, which is used to set the number of concurrently
	// consumed messages.Value range: 1-200
	SyncReplication bool `json:"sync_replication,omitempty"`
	// aging time in hours, value range: 1-168, defaults to 72
	RetentionTime int `json:"retention_time,omitempty"`
	// Topic configuration.
	TopicOtherConfigs []TopicOtherConfigs `json:"topic_other_configs,omitempty"`
	// Topic description.
	Description string `json:"topic_desc,omitempty"`
}

type TopicOtherConfigs struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Create a kafka topic with given parameters
func Create(client *golangsdk.ServiceClient, instanceID string, opts CreateOpts) (*CreateResponse, error) {
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
	Name string `json:"name"`
}
