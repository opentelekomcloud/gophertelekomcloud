package virtual_interface

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type ListOpts struct {
	ID string `q:"id"`
}

type VirtualInterface struct {
	ID                  string `json:"id"`
	TenantID            string `json:"tenant_id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	DirectConnectID     string `json:"direct_connect_id"`
	VgwID               string `json:"vgw_id"`
	Type                string `json:"type"`
	ServiceType         string `json:"service_type"`
	VLAN                int    `json:"vlan"`
	Bandwidth           int    `json:"bandwidth"`
	LocalGatewayV4IP    string `json:"local_gateway_v4_ip"`
	RemoteGatewayV4IP   string `json:"remote_gateway_v4_ip"`
	RouteMode           string `json:"route_mode"`
	BGPASN              int    `json:"bgp_asn"`
	BGPMD5              string `json:"bgp_md5"`
	RemoteEPGroupID     string `json:"remote_ep_group_id"`
	ServiceEPGroupID    string `json:"service_ep_group_id"`
	CreateTime          string `json:"create_time"`
	Status              string `json:"status"`
	AdminStateUp        bool   `json:"admin_state_up"`
	AddressFamily       string `json:"address_family"`
	EnableBFD           bool   `json:"enable_bfd"`
	HealthCheckSourceIP string `json:"health_check_source_ip"`
	RateLimit           bool   `json:"rate_limit"`
	RouteLimit          int    `json:"route_limit"`
	RegionID            string `json:"region_id"`
	EnableNQA           bool   `json:"enable_nqa"`
	EnableGRE           bool   `json:"enable_gre"`
	LocalGatewayV6IP    string `json:"local_gateway_v6_ip"`
	RemoteGatewayV6IP   string `json:"remote_gateway_v6_ip"`
	LocalGRETunnelIP    string `json:"local_gre_tunnel_ip"`
	RemoteGRETunnelIP   string `json:"remote_gre_tunnel_ip"`
	LagID               string `json:"lag_id"`
}

func List(c *golangsdk.ServiceClient, id string) ([]VirtualInterface, error) {
	// GET https://{Endpoint}/v2.0/{project_id}/virtual-interfaces?id={id}
	raw, err := c.Get(c.ServiceURL(fmt.Sprintf("dcaas/virtual-gateways?id=%s", id)), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []VirtualInterface
	err = extract.IntoSlicePtr(raw.Body, &res, "virtual_interfaces")
	return res, err
}
