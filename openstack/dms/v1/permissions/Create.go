package permissions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	Name     string         `json:"name" required:"true"`
	Policies []CreatePolicy `json:"policies" required:"true"`
}

type CreatePolicy struct {
	UserName     string `json:"user_name,omitempty"`
	AccessPolicy string `json:"access_policy,omitempty"`
}

func Create(client *golangsdk.ServiceClient, instanceId string, opts []CreateOpts) error {
	// PUT /v2/{project_id}/instances/{instance_id}/users
	b, err := build.RequestBody(opts, "topics")
	if err != nil {
		return err
	}
	url := client.ServiceURL("instances", instanceId, "topics", "accesspolicy")
	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
