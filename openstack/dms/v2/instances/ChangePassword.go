package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type PasswordOpts struct {
	NewPassword string `json:"new_password" required:"true"`
}

// ChangePassword is a method to update the password using given parameters.
func ChangePassword(client *golangsdk.ServiceClient, instanceId string, opts PasswordOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("instances", instanceId, "password"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
