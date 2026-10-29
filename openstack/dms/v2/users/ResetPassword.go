package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResetPasswordOpts struct {
	NewPassword string `json:"new_password,omitempty"`
}

// ResetPassword sends PUT /v2/{project_id}/instances/{instance_id}/users/{user_name}
func ResetPassword(client *golangsdk.ServiceClient, instanceId, userName string, opts ResetPasswordOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("instances", instanceId, "users", userName), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
