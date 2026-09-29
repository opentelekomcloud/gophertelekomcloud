package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// StatusResponse is returned by GetStatus.
type StatusResponse struct {
	// Maximum offset.
	MaxOffset int64 `json:"max_offset"`
	// Minimum offset.
	MinOffset int64 `json:"min_offset"`
	// Brokers of the topic.
	Brokers []StatusBroker `json:"brokers"`
}

type StatusBroker struct {
	// Node name.
	BrokerName string `json:"broker_name"`
	// Queue list.
	Queues []StatusQueue `json:"queues"`
}

type StatusQueue struct {
	// Queue ID.
	ID int `json:"id"`
	// Minimum offset.
	MinOffset int64 `json:"min_offset"`
	// Maximum offset.
	MaxOffset int64 `json:"max_offset"`
	// Time of the last message.
	LastMessageTime int64 `json:"last_message_time"`
}

// GetStatus queries the number of messages in a topic of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/topics/{topic}/status
func GetStatus(client *golangsdk.ServiceClient, instanceID, topic string) (*StatusResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "topics", topic, "status").Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res StatusResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
