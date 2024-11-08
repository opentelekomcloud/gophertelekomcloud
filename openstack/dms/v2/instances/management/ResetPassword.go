package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type PasswordOpts struct {
	NewPassword string `json:"new_password" required:"true"`
}

// ResetPassword is used to reset the password for an instance with SSL enabled.
// Send POST to /v2/{project_id}/instances/{instance_id}/password
func ResetPassword(client *golangsdk.ServiceClient, instanceId string, opts PasswordOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("instances", instanceId, "password"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
