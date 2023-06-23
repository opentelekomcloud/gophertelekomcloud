package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

func ResetPassword(client *golangsdk.ServiceClient, instanceId, userName, password string) error {
	// PUT /v2/{project_id}/instances/{instance_id}/users
	b, err := build.RequestBody(password, "new_password")
	if err != nil {
		return err
	}
	url := client.ServiceURL("instances", instanceId, "users", userName)
	_, err = client.Put(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
