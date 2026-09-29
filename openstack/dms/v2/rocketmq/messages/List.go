package messages

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Topic name.
	Topic string `q:"topic,required"`
	// Queue.
	Queue string `q:"queue,omitempty"`
	// Number of records to query.
	Limit int `q:"limit,omitempty"`
	// Offset where the query starts. Must be greater than or equal to 0.
	Offset int `q:"offset,omitempty"`
	// Message key.
	Key string `q:"key,omitempty"`
	// Start time, as the offset in milliseconds
	// from 1970-01-01 00:00:00 UTC to the specified time.
	// Mandatory when MsgID is not used for query.
	StartTime string `q:"start_time,omitempty"`
	// End time, as the offset in milliseconds
	// from 1970-01-01 00:00:00 UTC to the specified time.
	// Mandatory when MsgID is not used for query.
	EndTime string `q:"end_time,omitempty"`
	// Message ID. Mandatory when a time range is not used for query.
	MsgID string `q:"msg_id,omitempty"`
}

// ListResponse is returned by List.
type ListResponse struct {
	// Message list.
	Messages []Message `json:"messages"`
	// Total number of messages.
	Total int `json:"total"`
}

type Message struct {
	// Message ID.
	MsgID string `json:"msg_id"`
	// Instance ID.
	InstanceID string `json:"instance_id"`
	// Topic name.
	Topic string `json:"topic"`
	// Time when the message is stored.
	StoreTimestamp int64 `json:"store_timestamp"`
	// Time when the message is generated.
	BornTimestamp int64 `json:"born_timestamp"`
	// Number of retry times.
	ReconsumeTimes int `json:"reconsume_times"`
	// Message body.
	Body string `json:"body"`
	// Message body checksum.
	BodyCRC int64 `json:"body_crc"`
	// Storage size.
	StoreSize int64 `json:"store_size"`
	// Message attribute list.
	PropertyList []Property `json:"property_list"`
	// IP address of the host that generates the message.
	BornHost string `json:"born_host"`
	// IP address of the host that stores the message.
	StoreHost string `json:"store_host"`
	// Queue ID.
	QueueID int `json:"queue_id"`
	// Offset in the queue.
	QueueOffset int64 `json:"queue_offset"`
}

// List the messages of a RocketMQ instance.
// Send GET to /v2/{project_id}/rocketmq/instances/{instance_id}/messages
func List(client *golangsdk.ServiceClient, instanceID string, opts ListOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "messages").
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

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
