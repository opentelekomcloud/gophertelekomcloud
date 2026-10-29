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

// Create sends POST /v1/{project_id}/instances/{instance_id}/topics/accesspolicy
func Create(client *golangsdk.ServiceClient, instanceId string, opts []CreateOpts) error {
	b, err := build.RequestBody(opts, "topics")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("instances", instanceId, "topics", "accesspolicy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
