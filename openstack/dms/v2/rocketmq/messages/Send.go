package messages

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type SendOpts struct {
	// Topic name.
	Topic string `json:"topic" required:"true"`
	// Message content.
	Body string `json:"body" required:"true"`
	// Feature list.
	PropertyList []Property `json:"property_list,omitempty"`
}

// SendResponse is returned by Send.
type SendResponse struct {
	// Topic name.
	Topic string `json:"topic"`
	// Message content.
	Body string `json:"body"`
	// Feature list.
	PropertyList []Property `json:"property_list"`
	// Message ID.
	MsgID string `json:"msg_id"`
	// Queue ID.
	QueueID int `json:"queue_id"`
	// Queue offset.
	QueueOffset int64 `json:"queue_offset"`
	// Broker name.
	BrokerName string `json:"broker_name"`
}

// Send a message to a topic of a RocketMQ instance.
// Send POST to /v2/{project_id}/rocketmq/instances/{instance_id}/messages
func Send(client *golangsdk.ServiceClient, instanceID string, opts SendOpts) (*SendResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("rocketmq", "instances", instanceID, "messages").Build()
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

	var res SendResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
