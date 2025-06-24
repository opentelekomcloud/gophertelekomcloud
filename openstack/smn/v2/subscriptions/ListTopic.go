package subscriptions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListTopicOpts struct {
	TopicUrn string `json:"-"`
	Offset   string `q:"offset,omitempty"`
	Limit    int    `q:"limit,omitempty"`
}

// ListTopic subscriptions of specified topic
func ListTopic(client *golangsdk.ServiceClient, opts ListTopicOpts) ([]Subscription, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("topics", opts.TopicUrn, "subscriptions").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	// GET /v2/{project_id}/notifications/topics/{topic_urn}/subscriptions?offset={offset}&limit={limit}
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
