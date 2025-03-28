package servicegroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Protected object ID, which is used to distinguish between Internet border protection and VPC border protection
	// after a cloud firewall is created. You can obtain the ID by calling the Get function in management package.
	// In the return value, find the ID in ProtectObjects[n].ObjectID.
	// If the value of type is 0, the protected object ID belongs to the Internet border.
	// If the value of type is 1, the protected object ID belongs to the VPC border.
	ObjectID string `json:"object_id" required:"true"`
	// Service group name.
	Name string `json:"name" required:"true"`
	// Description.
	Description string `json:"description,omitempty"`
}

// This function is used to create a service group.
func CreateServiceGroup(client *golangsdk.ServiceClient, opts CreateOpts) (*ServiceSetId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/service-set
	raw, err := client.Post(client.ServiceURL("service-set"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res ServiceSetDataResponse
	return &res.Data, extract.Into(raw.Body, &res)
}
