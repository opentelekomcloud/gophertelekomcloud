package groups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get the details about a specified consumer group of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/groups/{group}
func Get(client *golangsdk.ServiceClient, instanceID, group string) (*ConsumerGroup, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "groups", group).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ConsumerGroup
	err = extract.Into(raw.Body, &res)
	return &res, err
}
