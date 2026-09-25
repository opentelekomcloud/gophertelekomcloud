package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListTopicsOpts struct {
	// Topic to be queried. If not specified, the topic list of the consumer
	// group is returned, otherwise the consumption details of the topic.
	Topic string `q:"topic,omitempty"`
	// Maximum number of topics that can be returned in a query.
	// Range: 1-50. Default: 10.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
}

// ListTopicsResponse is returned by ListTopics.
type ListTopicsResponse struct {
	// Topic list. Returned only when the topic list is queried.
	Topics []string `json:"topics"`
	// Total number of topics. Returned only when the topic list is queried.
	Total int `json:"total"`
	// Total number of accumulated messages.
	Lag int64 `json:"lag"`
	// Total number of messages.
	MaxOffset int64 `json:"max_offset"`
	// Number of consumed messages.
	ConsumerOffset int64 `json:"consumer_offset"`
	// Associated brokers of the topic. Returned only when the topic details are queried.
	Brokers []TopicBroker `json:"brokers"`
}

type TopicBroker struct {
	// Name of the associated broker.
	BrokerName string `json:"broker_name"`
	// Queue details of the associated broker.
	Queues []TopicQueue `json:"queues"`
}

type TopicQueue struct {
	// Queue ID.
	ID int `json:"id"`
	// Total number of accumulated messages in the queue.
	Lag int64 `json:"lag"`
	// Total number of messages in the queue.
	BrokerOffset int64 `json:"broker_offset"`
	// Number of consumed messages.
	ConsumerOffset int64 `json:"consumer_offset"`
	// Time (UNIX, in milliseconds) when the latest consumed message was stored.
	LastMessageTime int64 `json:"last_message_time"`
}

// ListTopics queries the topic list of a consumer group or,
// when opts.Topic is set, the consumption details of that topic.
// Send GET to /v2/{project_id}/instances/{instance_id}/groups/{group}/topics
func ListTopics(client *golangsdk.ServiceClient, instanceID, group string, opts ListTopicsOpts) (*ListTopicsResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", instanceID, "groups", group, "topics").
		WithQueryParams(&opts).
		Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListTopicsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
