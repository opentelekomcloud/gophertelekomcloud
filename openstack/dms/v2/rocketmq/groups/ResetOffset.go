package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ResetOffsetOpts struct {
	// Topic to be reset.
	Topic string `json:"topic" required:"true"`
	// Reset time, as the offset in milliseconds
	// from 1970-01-01 00:00:00 UTC to the specified time.
	Timestamp int64 `json:"timestamp"`
}

// ResetOffsetResponse is returned by ResetOffset.
type ResetOffsetResponse struct {
	// Queues to be reset.
	Queues []ResetQueue `json:"queues"`
}

type ResetQueue struct {
	// Broker where the queue is located.
	BrokerName string `json:"broker_name"`
	// Queue ID.
	QueueID int `json:"queue_id"`
	// Reset consumer offset.
	TimestampOffset int64 `json:"timestamp_offset"`
}

// ResetOffset resets the consumer offset of a consumer group for a topic.
// Send POST to /v2/{project_id}/rocketmq/instances/{instance_id}/groups/{group}/reset-message-offset
func ResetOffset(client *golangsdk.ServiceClient, instanceID, group string, opts ResetOffsetOpts) (*ResetOffsetResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "groups", group, "reset-message-offset").
		Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ResetOffsetResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
