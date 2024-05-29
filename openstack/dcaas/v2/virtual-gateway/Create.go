package virtual_gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the virtual gateway name.
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the virtual gateway.
	Description string `json:"description,omitempty"`
	// Specifies the ID of the VPC to be accessed.
	VpcId string `json:"vpc_id" required:"true"`
	// Specifies the ID of the local endpoint group that records IPV4 CIDR blocks of the VPC subnets.
	// Either `LocalEndpointGroupId` or `LocalEndpointGroupIpv6Id` have to be specified
	LocalEndpointGroupId string `json:"local_ep_group_id,omitempty"`
	// Specifies the ID of the local endpoint group that records IPV6 CIDR blocks of the VPC subnets.
	// Either `LocalEndpointGroupId` or `LocalEndpointGroupIpv6Id` have to be specified
	LocalEndpointGroupIpv6Id string `json:"local_ep_group_ipv6_id,omitempty"`
	// Specifies the BGP ASN of the virtual gateway.
	BgpAsn int `json:"bgp_asn,omitempty"`
	// Specifies the ID of the physical device used by the virtual gateway.
	DeviceId string `json:"device_id,omitempty"`
	// Specifies the ID of the redundant physical device used by the virtual gateway.
	RedundantDeviceId string `json:"redundant_device_id,omitempty"`
	// Specifies the virtual gateway type. The value can only be default.
	Type string `json:"type"`
	// Specifies the project ID.
	ProjectId string `json:"tenant_id,omitempty"`
	// Specifies the administrative status of the virtual gateway.
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
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
