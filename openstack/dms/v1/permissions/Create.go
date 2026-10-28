package permissions

import (
	"fmt"

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
	if instanceId == "" {
		return golangsdk.ErrMissingInput{Argument: "instanceId"}
	}
	if len(opts) == 0 {
		return golangsdk.ErrMissingInput{Argument: "topics"}
	}
	for i, topic := range opts {
		if err := build.ValidateTags(topic); err != nil {
			return fmt.Errorf("topics[%d]: %w", i, err)
		}
	}

	b, err := build.RequestBody(opts, "topics")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("instances", instanceId, "topics", "accesspolicy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
