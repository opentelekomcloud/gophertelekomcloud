package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	UserName   string `json:"user_name,omitempty"`
	UserPasswd string `json:"user_passwd,omitempty"`
}

func Create(client *golangsdk.ServiceClient, instanceId string, opts CreateOpts) error {
	// PUT /v2/{project_id}/instances/{instance_id}/users
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	url := client.ServiceURL("instances", instanceId, "users")
	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
