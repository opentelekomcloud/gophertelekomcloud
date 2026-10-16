package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get a topic with detailed information by instance id and topic name
// Send GET /v2/{project_id}/instances/{instance_id}/management/topics/{topic}
func Get(client *golangsdk.ServiceClient, instanceID, topicName string) (*TopicDetail, error) {
	if instanceID == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "instanceID"}
	}
	if topicName == "" {
		return nil, golangsdk.ErrMissingInput{Argument: "topicName"}
	}

	raw, err := client.Get(client.ServiceURL("instances", instanceID, "management", "topics", topicName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res TopicDetail
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// TopicDetail includes the detail parameters of an topic
type TopicDetail struct {
	// Topic name.
	Name string `json:"topic"`
	// Partition list.
	Partitions []Partition `json:"partitions"`
	// List of consumer groups that subscribe to the topic.
	GroupSubscribed []string `json:"group_subscribed"`
}

// Partition represents the details of a partition
type Partition struct {
	// Partition ID.
	Partition int `json:"partition"`
	// Replica list.
	Replicas []Replica `json:"replicas"`
	// ID of the broker where the leader replica resides.
	Leader int `json:"leader"`
	// LEO of the partition leader replica.
	Leo int `json:"leo"`
	// High watermark (HW) of the partition.
	Hw int `json:"hw"`
	// Log start offset (LSO) of the partition leader replica.
	Lso int `json:"lso"`
	// Time when the last message was written to the partition.
	// The value is a Unix timestamp in milliseconds.
	UpdateTimestamp int64 `json:"last_update_timestamp"`
}

// Replica represents the details of a replica
type Replica struct {
	// ID of the broker where the replica resides.
	Broker int `json:"broker"`
	// Whether the replica is the leader.
	Leader bool `json:"leader"`
	// Whether the replica is in the ISR.
	InSync bool `json:"in_sync"`
	// Current log size of the replica. Unit: byte.
	Size int `json:"size"`
	// Number of messages that lag behind the high watermark in the replica.
	Lag int64 `json:"lag"`
}
