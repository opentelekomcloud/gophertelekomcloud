package addressgroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// IP address group name.
	Name string `json:"name,omitempty"`
	// Description.
	Description string `json:"description,omitempty"`
}

// This function is used to update address group information.
// groupId: Address group ID. It is the same as ID retuned while creating an address group.
func UpdateAddressGroup(client *golangsdk.ServiceClient, groupId string, opts UpdateOpts) (*AddressSetId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/address-sets/{set_id}
	raw, err := client.Put(client.ServiceURL("address-sets", groupId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res UpdateResponse
	return &res.Data, extract.Into(raw.Body, &res)
}

type UpdateResponse struct {
	Data AddressSetId `json:"data"`
}
