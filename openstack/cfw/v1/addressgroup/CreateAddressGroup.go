package addressgroup

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
	// IP address group name.
	Name string `json:"name" required:"true"`
	// Description.
	Description string `json:"description,omitempty"`
	// Internet protocol type of an address: 0 (IPv4), 1 (IPv6).
	AddressType *int `json:"address_type,omitempty"`
}

// This function is used to add an address group.
func CreateAddressGroup(client *golangsdk.ServiceClient, opts CreateOpts) (*AddressSetId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/address-set
	raw, err := client.Post(client.ServiceURL("address-set"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res CreateResponse
	return &res.Data, extract.Into(raw.Body, &res)
}

type CreateResponse struct {
	Data AddressSetId `json:"data"`
}
