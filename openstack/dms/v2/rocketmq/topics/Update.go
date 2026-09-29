package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Topic description.
	Description *string `json:"topic_desc,omitempty"`
}

// Update a specified topic of a RocketMQ instance.
// Send PUT to /v2/{project_id}/instances/{instance_id}/topics/{topic}
func Update(client *golangsdk.ServiceClient, instanceID, topic string, opts UpdateOpts) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "topics", topic).Build()
	if err != nil {
		return err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
