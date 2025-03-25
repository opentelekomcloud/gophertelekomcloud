package addressgroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query address group details.
// groupId: Address group ID. It is the same as ID retuned while creating an address group.
func GetAddressGroup(client *golangsdk.ServiceClient, groupId string) (*AddressGroupData, error) {
	// GET /v1/{project_id}/address-sets/{set_id}
	raw, err := client.Get(client.ServiceURL("address-sets", groupId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResponse
	err = extract.Into(raw.Body, &res)
	return &res.Data, err
}

type GetResponse struct {
	Data AddressGroupData `json:"data"`
}

type AddressGroupData struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	AddressSetType int    `json:"address_set_type"`
	AddressType    int    `json:"address_type"`
}
