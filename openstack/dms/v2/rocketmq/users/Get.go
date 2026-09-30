package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get the details about a specified user of a RocketMQ instance.
// Send GET to /v2/{project_id}/instances/{instance_id}/users/{user_name}
func Get(client *golangsdk.ServiceClient, instanceID, userName string) (*User, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "users", userName).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res User
	err = extract.Into(raw.Body, &res)
	return &res, err
}
