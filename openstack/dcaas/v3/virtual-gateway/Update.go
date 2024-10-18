package virtual_gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// The list of IPv4 subnets from the virtual gateway to access cloud services, which is usually the CIDR block of
	// the VPC.
	LocalEpGroup []string `json:"local_ep_group,omitempty"`
	// The list of IPv6 subnets from the virtual gateway to access cloud services, which is usually the CIDR block of
	// the VPC.
	LocalEpGroupIpv6 []string `json:"local_ep_group_ipv6,omitempty"`
	// Specifies the name of the virtual gateway.
	// The valid length is limited from 0 to 64, only chinese and english letters, digits, hyphens (-), underscores (_)
	// and dots (.) are allowed.
	// The name must start with a chinese or english letter, and the Chinese characters must be in **UTF-8** or
	// **Unicode** format.
	Name string `json:"name,omitempty"`
	// Specifies the description of the virtual gateway.
	// The description contain a maximum of 64 characters and the angle brackets (< and >) are not allowed.
	// Chinese characters must be in **UTF-8** or **Unicode** format.
	Description *string `json:"description,omitempty"`
}

func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (*VirtualGateway, error) {
	b, err := build.RequestBody(opts, "virtual_gateway")
	if err != nil {
		return nil, err
	}
	raw, err := c.Put(c.ServiceURL("dcaas", "virtual-gateways", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	var res VirtualGateway
	err = extract.IntoStructPtr(raw.Body, &res, "virtual_gateway")
	return &res, err
}
