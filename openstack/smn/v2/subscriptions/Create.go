package subscriptions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Message endpoint
	Endpoint string `json:"endpoint" required:"true"`
	// Protocol of the message endpoint
	Protocol string `json:"protocol" required:"true"`
	// Description of the subscription
	Remark string `json:"remark,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts, topicUrn string) (*CreateResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/notifications/topics/{topic_urn}/subscriptions
	raw, err := client.Post(client.ServiceURL("topics", topicUrn, "subscriptions"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateResponse struct {
	RequestID       string `json:"request_id"`
	SubscriptionUrn string `json:"subscription_urn"`
}
