package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get the details about a specified RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}
func Get(client *golangsdk.ServiceClient, id string) (*Instance, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", id).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Instance
	err = extract.Into(raw.Body, &res)
	return &res, err
}
