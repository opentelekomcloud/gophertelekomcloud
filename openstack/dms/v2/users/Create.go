package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	UserName   string `json:"user_name,omitempty"`
	UserDesc   string `json:"user_desc,omitempty"`
	UserPasswd string `json:"user_passwd,omitempty"`
}

// Create sends POST /v2/{project_id}/instances/{instance_id}/users
func Create(client *golangsdk.ServiceClient, instanceId string, opts CreateOpts) error {
	if instanceId == "" {
		return golangsdk.ErrMissingInput{Argument: "instanceId"}
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("instances", instanceId, "users"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
