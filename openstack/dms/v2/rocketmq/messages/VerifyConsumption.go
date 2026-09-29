package messages

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type VerifyConsumptionOpts struct {
	// Consumer group name.
	Group string `json:"group" required:"true"`
	// Topic to which the messages belong.
	Topic string `json:"topic,omitempty"`
	// ID of the consumer client.
	ClientID string `json:"client_id" required:"true"`
	// IDs of the messages to be consumed again.
	MsgIDList []string `json:"msg_id_list,omitempty"`
}

// VerifyConsumption makes a consumer client consume the specified messages again
// to verify its consumption.
// Send POST to /v2/{project_id}/rocketmq/instances/{instance_id}/messages/resend
func VerifyConsumption(client *golangsdk.ServiceClient, instanceID string, opts VerifyConsumptionOpts) (*ResendResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("rocketmq", "instances", instanceID, "messages", "resend").
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
