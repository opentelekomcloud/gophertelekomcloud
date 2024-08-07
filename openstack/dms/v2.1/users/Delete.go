package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	Action string   `json:"action,omitempty"`
	Users  []string `json:"users,omitempty"`
}

func Delete(client *golangsdk.ServiceClient, instanceId string, opts DeleteOpts) error {
	// POST /v2/{project_id}/instances/{instance_id}/users
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}
	url := client.ServiceURL("instances", instanceId, "users")
	_, err = client.Put(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
