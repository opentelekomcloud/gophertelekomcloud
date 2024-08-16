package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

const passwordPath = "autotopic"

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

	_, err = client.Post(client.ServiceURL(instances.ResourcePath, instanceId, passwordPath), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return err
}
