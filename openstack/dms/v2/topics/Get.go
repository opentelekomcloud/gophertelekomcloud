package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get a topic with detailed information by instance id and topic name
func Get(client *golangsdk.ServiceClient, instanceID, topicName string) (*TopicDetail, error) {
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
	Name            string      `json:"topic"`
	Partitions      []Partition `json:"partitions"`
	GroupSubscribed []string    `json:"group_subscribed"`
}

// Partition represents the details of a partition
type Partition struct {
	Partition int       `json:"partition"`
	Replicas  []Replica `json:"replicas"`
	// Node ID
	Leader int `json:"leader"`
	// Log End Offset
	Leo int `json:"leo"`
	// High Watermark
	Hw int `json:"hw"`
	// Log Start Offset
	Lso int `json:"lso"`
	// time stamp
	UpdateTimestamp int64 `json:"last_update_timestamp"`
}

// Replica represents the details of a replica
type Replica struct {
	Broker int   `json:"broker"`
	Leader bool  `json:"leader"`
	InSync bool  `json:"in_sync"`
	Size   int   `json:"size"`
	Lag    int64 `json:"lag"`
}
