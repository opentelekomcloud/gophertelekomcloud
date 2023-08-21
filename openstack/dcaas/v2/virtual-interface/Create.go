package virtual_interface

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	TenantID          string `json:"tenant_id,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	DirectConnectID   string `json:"direct_connect_id"  required:"true"`
	VgwID             string `json:"vgw_id" required:"true"`
	Type              string `json:"type" required:"true"`
	ServiceType       string `json:"service_type" required:"true"`
	VLAN              int    `json:"vlan" required:"true"`
	Bandwidth         int    `json:"bandwidth" required:"true"`
	LocalGatewayV4IP  string `json:"local_gateway_v4_ip" required:"true"`
	RemoteGatewayV4IP string `json:"remote_gateway_v4_ip" required:"true"`
	RouteMode         string `json:"route_mode" required:"true"`
	BGPASN            int    `json:"bgp_asn,omitempty"`
	BGPMD5            string `json:"bgp_md5,omitempty"`
	RemoteEPGroupID   string `json:"remote_ep_group_id" required:"true"`
	AdminStateUp      bool   `json:"admin_state_up,omitempty"`
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
