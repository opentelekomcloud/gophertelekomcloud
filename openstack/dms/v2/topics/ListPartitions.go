package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListPartitionsOpts represents the query parameters of the partition list request
type ListPartitionsOpts struct {
	// Offset. The records after this offset will be queried.
	Offset *int `q:"offset,omitempty"`
	// Maximum number of records that can be returned.
	Limit *int `q:"limit,omitempty"`
}

// ListPartitions queries the partition list of a topic
// Send GET /v2/{project_id}/kafka/instances/{instance_id}/topics/{topic}/partitions
func ListPartitions(client *golangsdk.ServiceClient, instanceID, topicName string, opts ListPartitionsOpts) (*ListPartitionsResponse, error) {
	if instanceID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "instanceID"}
	}
	if topicName == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "topicName"}
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("kafka", "instances", instanceID, "topics", topicName, "partitions").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListPartitionsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// ListPartitionsResponse is a struct that contains the partition list response
type ListPartitionsResponse struct {
	// Total records.
	Total int `json:"total"`
	// Partition array.
	Partitions []TopicPartition `json:"partitions"`
}

// TopicPartition represents one partition in the partition list of a topic
type TopicPartition struct {
	// Partition ID.
	Partition int `json:"partition"`
	// Start offset.
	StartOffset int64 `json:"start_offset"`
	// Last offset.
	LastOffset int64 `json:"last_offset"`
	// Number of messages in a partition.
	MessageCount int64 `json:"message_count"`
	// Last update time.
	LastUpdateTime int64 `json:"last_update_time"`
}
