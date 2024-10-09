package virtual_gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	// The ID of the VPC connected to the virtual gateway.
	VpcId string `json:"vpc_id" required:"true"`
	// The list of IPv4 subnets from the virtual gateway to access cloud services, which is usually the CIDR block of
	// the VPC.
	LocalEpGroup []string `json:"local_ep_group" required:"true"`
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
	Description string `json:"description,omitempty"`
	// The local BGP ASN of the virtual gateway.
	BgpAsn int `json:"bgp_asn,omitempty"`
	// The key/value pairs to associate with the virtual gateway.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*VirtualGateway, error) {
	b, err := build.RequestBody(opts, "virtual_gateway")
	if err != nil {
		return nil, err
	}
	raw, err := c.Post(c.ServiceURL("dcaas", "virtual-gateways"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res VirtualGateway
	err = extract.IntoStructPtr(raw.Body, &res, "virtual_gateway")
	return &res, err
}
