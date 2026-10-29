package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteOpts struct {
	Action string   `json:"action,omitempty"`
	Users  []string `json:"users,omitempty"`
}

// Delete sends PUT /v2/{project_id}/instances/{instance_id}/users
func Delete(client *golangsdk.ServiceClient, instanceId string, opts DeleteOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL("instances", instanceId, "users"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
