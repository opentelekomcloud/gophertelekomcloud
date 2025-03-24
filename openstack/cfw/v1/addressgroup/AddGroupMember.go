package addressgroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type AddGroupMemberOpts struct {
	// Address group ID
	SetId string `json:"set_id,omitempty"`
	// Address group member list.
	AddressItems []AddressItem `json:"address_items,omitempty"`
}

type AddressItem struct {
	// Internet protocol type of an address: 0 (IPv4), 1 (IPv6).
	AddressType *int `json:"address_type,omitempty"`
	// IP address.
	Address string `json:"address" required:"true"`
	// Description.
	Description string `json:"description,omitempty"`
}

// This function is used to add a member to an address group.
func AddGroupMember(client *golangsdk.ServiceClient, opts AddGroupMemberOpts) (*AddressItemsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/address-items
	raw, err := client.Post(client.ServiceURL("address-items"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res AddGroupMemberResponse
	return &res.Data, extract.Into(raw.Body, &res)
}

type AddGroupMemberResponse struct {
	Data AddressItemsResponse `json:"data"`
}

type AddressItemsResponse struct {
	// List of address group member IDs.
	Items []AddressItemId `json:"items"`
	// List of covered IP addresses.
	CoveredIp []CoveredIPVO `json:"covered_ip"`
}

type AddressItemId struct {
	// ID of an address group member.
	Id string `json:"id"`
}

type CoveredIPVO struct {
	// IP address
	IP string `json:"ip"`
	// Cover an IP address.
	CoveredIP string `json:"covered_Ip"`
}
