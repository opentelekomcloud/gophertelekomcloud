package addressgroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain details about an IP address group.
func Get(client *golangsdk.ServiceClient, groupId string) (*AddressGroup, error) {
	// GET /v3/{project_id}/vpc/address-groups/{address_group_id}
	raw, err := client.Get(client.ServiceURL("address-groups", groupId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res AddressGroupResponse
	err = extract.Into(raw.Body, &res)
	return &res.AddressGroup, err
}
