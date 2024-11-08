package protectedinstances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular Instance based on its unique ID.
func Get(client *golangsdk.ServiceClient, instanceId string) (*Instance, error) {
	// GET /v1/{project_id}/protected-instances/{protected_instance_id}
	raw, err := client.Get(client.ServiceURL("protected-instances", instanceId), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res.ProtectedInstance, err

}

type GetResponse struct {
	// Specifies the details about a protected instance.
	ProtectedInstance Instance `json:"protected_instance"`
}
