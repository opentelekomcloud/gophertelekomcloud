package blackwhitelist

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
	// Blacklist/Whitelist type: 4 (blacklist), 5 (whitelist).
	ListType int `json:"list_type" required:"true"`
	// Address direction: 0 (source), 1 (destination).
	Direction *int `json:"direction" required:"true"`
	// Internet protocol type of an address: 0 (IPv4), 1 (IPv6).
	AddressType *int `json:"address_type" required:"true"`
	// IP address.
	Address string `json:"address" required:"true"`
	// Protocol type: 6 (TCP), 17 (UDP), 1 (ICMP), 58 (ICMPv6), or -1 (any).
	// Cannot be left blank when type is set to 0 (manual) and can be omitted when type is set to 1 (automatic).
	Protocol int `json:"protocol" required:"true"`
	// Destination port.
	Port string `json:"port" required:"true"`
	// Description.
	Description string `json:"description,omitempty"`
}

// This function is used to create a blacklist or whitelist rule.
func CreateBlacklistOrWhitelistRule(client *golangsdk.ServiceClient, firewallId string, opts CreateOpts) (*BlackWhiteListId, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/black-white-list
	raw, err := client.Post(client.ServiceURL("black-white-list"), b, nil, &golangsdk.RequestOpts{
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
	Data BlackWhiteListId `json:"data"`
}
