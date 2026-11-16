package notifications

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	TopicURN string `json:"topic_urn" required:"true"`
	// TopicScene accepts SCALING_UP, SCALING_UP_FAIL, SCALING_DOWN,
	// SCALING_DOWN_FAIL, and SCALING_GROUP_ABNORMAL.
	TopicScene []string `json:"topic_scene" required:"true"`
}

// Create configures a notification topic and its scaling events for an AS group.
func Create(client *golangsdk.ServiceClient, opts CreateOpts, scalingGroupID string) (*Notification, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("scaling_notification", scalingGroupID).Build()
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{OkCodes: []int{200}})
	if err != nil {
		return nil, err
	}
	var res Notification
	if err = extract.Into(raw.Body, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type Notification struct {
	TopicURN   string   `json:"topic_urn"`
	TopicScene []string `json:"topic_scene"`
	TopicName  string   `json:"topic_name"`
}
