package blackwhitelist

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Address direction: 0 (source), 1 (destination).
	Direction *int `json:"direction,omitempty"`
	// Internet protocol type of an address: 0 (IPv4), 1 (IPv6).
	AddressType *int `json:"address_type,omitempty"`
	// IP address.
	Address string `json:"address" required:"true"`
	// Protocol type: 6 (TCP), 17 (UDP), 1 (ICMP), 58 (ICMPv6), or -1 (any).
	// Cannot be left blank when type is set to 0 (manual) and can be omitted when type is set to 1 (automatic).
	Protocol int `json:"protocol,omitempty"`
	// Destination port.
	Port string `json:"port,omitempty"`
	// Description.
	Description string `json:"description,omitempty"`
}

// This function is used to update a blacklist or whitelist.
func UpdateBlacklistOrWhitelistRule(client *golangsdk.ServiceClient, listId string, opts UpdateOpts) (*BlackWhiteListId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/black-white-list/{list_id}
	raw, err := client.Put(client.ServiceURL("black-white-list", listId), b, nil, &golangsdk.RequestOpts{
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
	Data BlackWhiteListId `json:"data"`
}
