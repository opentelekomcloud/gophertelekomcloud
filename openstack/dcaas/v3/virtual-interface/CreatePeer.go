package virtual_interface

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreatePeerOpts struct {
	// Specifies the name of the virtual interface peer.
	Name string `json:"name" required:"true"`
	// Provides supplementary information about the virtual interface peer.
	Description string `json:"description,omitempty"`
	// Specifies the gateway address of the virtual interface peer used on the cloud.
	LocalGatewayIP string `json:"local_gateway_ip" required:"true"`
	// Specifies the IPv4 interface address of the gateway on the on-premises network.
	// This parameter is mandatory if address_family is set to an IPv4 address.
	RemoteGatewayIP string `json:"remote_gateway_ip" required:"true"`
	// Specifies the address family of the virtual interface.
	// The value can be IPv4 or IPv6.
	AddressFamily string `json:"address_family" required:"true"`
	// Specifies the routing mode. The value can be static or bgp.
	RouteMode string `json:"route_mode" required:"true"`
	// Specifies the ASN of the BGP peer on the customer side.
	BGPASN int `json:"bgp_asn,omitempty"`
	// Specifies the MD5 password of the BGP peer.
	BGPMD5 string `json:"bgp_md5,omitempty"`
	// Specifies the remote subnet list, which records the CIDR blocks used in the on-premises data center.
	RemoteEpGroup []string `json:"remote_ep_group,omitempty"`
	// Specifies the ID of the virtual interface corresponding to the virtual interface peer.
	VifId string `json:"vif_id" required:"true"`
}

func CreatePeer(c *golangsdk.ServiceClient, opts CreatePeerOpts) (*VifPeer, error) {
	b, err := build.RequestBody(opts, "vif_peer")
	if err != nil {
		return nil, err
	}
	raw, err := c.Post(c.ServiceURL("dcaas", "vif-peers"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res VifPeer
	err = extract.IntoStructPtr(raw.Body, &res, "vif_peer")
	return &res, err
}
