package virtual_interface

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the project ID of another tenant, which is used to create virtual interfaces across tenants.
	ResourceTenantID string `json:"resource_tenant_id,omitempty"`
	// Specifies the virtual interface name.
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the virtual interface.
	Description string `json:"description,omitempty"`
	// Specifies the ID of the connection associated with the virtual interface.
	// When creating a virtual interface, you need to specify direct_connect_id or lag_id.
	// This parameter is mandatory when LAG is not supported at the site.
	DirectConnectID string `json:"direct_connect_id,omitempty"`
	// Specifies the type of the virtual interface. The value is private.
	Type string `json:"type" required:"true"`
	// Specifies the type of the access gateway. You do not need to set this parameter.
	ServiceType string `json:"service_type,omitempty"`
	// Specifies the customer VLAN to be connected.
	// If you select a hosted connection, the VLAN must be the same as that of the hosted connection.
	VLAN int `json:"vlan" required:"true"`
	// Specifies the virtual interface bandwidth.
	Bandwidth int `json:"bandwidth" required:"true"`
	// Specifies the IPv4 interface address of the gateway used on the cloud.
	// This parameter is mandatory if address_family is set to an IPv4 address.
	LocalGatewayV4IP string `json:"local_gateway_v4_ip,omitempty"`
	// Specifies the IPv4 interface address of the gateway on the on-premises network.
	// This parameter is mandatory if address_family is set to an IPv4 address.
	RemoteGatewayV4IP string `json:"remote_gateway_v4_ip,omitempty"`
	// Specifies the address family of the virtual interface.
	// The value can be IPv4 or IPv6.
	AddressFamily string `json:"address_family,omitempty"`
	// Specifies the IPv6 interface address of the gateway used on the cloud.
	// This parameter is mandatory if address_family is set to an IPv6 address.
	LocalGatewayV6IP string `json:"local_gateway_v6_ip,omitempty"`
	// Specifies the IPv6 interface address of the gateway on the on-premises network.
	// This parameter is mandatory if address_family is set to an IPv6 address.
	RemoteGatewayV6IP string `json:"remote_gateway_v6_ip,omitempty"`
	// Specifies the ID of the virtual gateway connected by the virtual interface.
	VgwId string `json:"vgw_id,omitempty" required:"true"`
	// Specifies the routing mode. The value can be static or bgp.
	RouteMode string `json:"route_mode" required:"true"`
	// Specifies the ASN of the BGP peer on the customer side.
	BGPASN int `json:"bgp_asn,omitempty"`
	// Specifies the MD5 password of the BGP peer.
	BGPMD5 string `json:"bgp_md5,omitempty"`
	// Specifies the remote subnet list, which records the CIDR blocks used in the on-premises data center.
	RemoteEpGroup []string `json:"remote_ep_group" required:"true"`
	// Specifies the subnets that access Internet services through a connection.
	ServiceEpGroup []string `json:"service_ep_group,omitempty"`
	// Specifies whether to enable BFD.
	EnableBfd bool `json:"enable_bfd,omitempty"`
	// Specifies whether to enable NQA.
	EnableNqa bool `json:"enable_nqa,omitempty"`
	// Specifies the ID of the LAG associated with the virtual interface.
	LagId string `json:"lag_id,omitempty"`
}

func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*VirtualInterface, error) {
	b, err := build.RequestBody(opts, "virtual_interface")
	if err != nil {
		return nil, err
	}
	raw, err := c.Post(c.ServiceURL("dcaas", "virtual-interfaces"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res VirtualInterface
	err = extract.IntoStructPtr(raw.Body, &res, "virtual_interface")
	return &res, err
}
