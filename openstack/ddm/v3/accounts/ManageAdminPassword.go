package accounts

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ManageAdminPassOpts struct {
	// Username of the administrator. The username:
	// Can include 1 to 32 characters.
	// Must start with a letter.
	// Can contain only letters, digits, and underscores (_).
	Name string `json:"name" required:"true"`
	// Password of the administrator. The password:
	// Can include 8 to 32 characters.
	// Must be a combination of uppercase letters, lowercase letters, digits, and the following special characters: ~!@#%^*-_=+?
	// Must be a strong password to improve security and prevent security risks such as brute force cracking.
	Password string `json:"password" required:"true"`
}

// This function is used to manage the password of the DDM instance administrator.
// If it is the first time to call this API, it is used to create an administrator and reset its password for a DDM instance.
// Then this API can only be used to update the administrator password.
func ManageAdminPass(client *golangsdk.ServiceClient, instanceId string, opts ManageAdminPassOpts) (*ManageAdminPassResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/instances/{instance_id}/admin-user
	raw, err := client.Put(client.ServiceURL("instances", instanceId, "admin-user"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ManageAdminPassResponse
	return &res, extract.Into(raw.Body, &res)
}

type ManageAdminPassResponse struct {
	// DDM instance ID
	InstanceId string `json:"instance_id"`
	// Task ID
	JobId string `json:"job_id"`
}
