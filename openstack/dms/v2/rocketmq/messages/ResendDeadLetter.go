package messages

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ResendDeadLetterOpts struct {
	// Dead letter topic name, e.g. %DLQ%{group}.
	Topic string `json:"topic,omitempty"`
	// IDs of the messages to be resent.
	MsgIDList []string `json:"msg_id_list,omitempty"`
}

// ResendDeadLetter resends dead letter messages of a RocketMQ instance.
// Send POST to /v2/{project_id}/rocketmq/instances/{instance_id}/messages/deadletter-resend
func ResendDeadLetter(client *golangsdk.ServiceClient, instanceID string, opts ResendDeadLetterOpts) (*ResendResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "messages", "deadletter-resend").
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

	var res ResendResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
