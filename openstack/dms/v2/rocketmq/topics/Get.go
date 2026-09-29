package topics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get the details about a specified topic of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/topics/{topic}
func Get(client *golangsdk.ServiceClient, instanceID, topic string) (*Topic, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "topics", topic).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Topic
	err = extract.Into(raw.Body, &res)
	return &res, err
}
