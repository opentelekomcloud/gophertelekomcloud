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
	// The following protocols are supported:
	// email: The endpoints are email address.
	// sms: The endpoints are phone numbers.
	// http and https: The endpoints are URLs.
	Protocol string `json:"protocol" required:"true"`
	// Description of the subscription
	Remark string `json:"remark,omitempty"`
	// Extended information
	Extension *SubscriptionExtension `json:"extension,omitempty"`
	// This parameter is mandatory when subscriptions need to be created in batches.
	// SMN allows you to create a maximum of 50 subscriptions at a time.
	Subscriptions []AddSubscriptions `json:"subscriptions,omitempty"`
}

type AddSubscriptions struct {
	// Message endpoint
	Endpoint string `json:"endpoint" required:"true"`
	// Protocol of the message endpoint
	// The following protocols are supported:
	// email: The endpoints are email address.
	// sms: The endpoints are phone numbers.
	// http and https: The endpoints are URLs.
	Protocol string `json:"protocol" required:"true"`
	// Description of the subscription
	Remark string `json:"remark,omitempty"`
	// Extended information
	Extension *SubscriptionExtension `json:"extension,omitempty"`
}

type SubscriptionExtension struct {
	// This is an HTTP header field, which can be customized within the field range.
	// The field content exists in the form of key/value pairs.
	// When a topic is used to send messages, confirmed subscription messages carry the user-defined HTTP header.
	Header map[string]string `json:"header,omitempty"`
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
