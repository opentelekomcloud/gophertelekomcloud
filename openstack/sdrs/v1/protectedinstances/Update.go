package protectedinstances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts contains all the values needed to update an Instance.
type UpdateOpts struct {
	// Instance name. Specifies the name of a protected instance.
	// The name can contain a maximum of 64 bytes.
	// The value can contain only letters (a to z and A to Z), digits (0 to 9), decimal points (.), underscores (_), and hyphens (-).
	Name string `json:"name" required:"true"`
}

// Update accepts a UpdateOpts struct and uses the values to update an Instance.The response code from api is 200
func Update(client *golangsdk.ServiceClient, instanceId string, opts UpdateOpts) (*Instance, error) {
	b, err := build.RequestBody(opts, "protected_instance")
	if err != nil {
		return nil, err
	}
	// PUT /v1/{project_id}/protected-instances/{protected_instance_id}
	raw, err := client.Put(client.ServiceURL("protected-instances", instanceId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateResponse
	err = extract.Into(raw.Body, &res)
	return &res.ProtectedInstance, err
}

type UpdateResponse struct {
	// Specifies the details about a protected instance.
	ProtectedInstance Instance `json:"protected_instance"`
}
