package subscriptions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Offset string `q:"offset,omitempty"`
	Limit  int    `q:"limit,omitempty"`
}

// List all the subscriptions
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Subscription, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("subscriptions").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	// GET /v2/{project_id}/notifications/subscriptions
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []Subscription
	err = extract.IntoSlicePtr(raw.Body, &res, "subscriptions")
	return res, err
}

type Subscription struct {
	TopicUrn        string `json:"topic_urn"`
	Protocol        string `json:"protocol"`
	SubscriptionUrn string `json:"subscription_urn"`
	Owner           string `json:"owner"`
	Endpoint        string `json:"endpoint"`
	Remark          string `json:"remark"`
	Status          int    `json:"status"`
}
