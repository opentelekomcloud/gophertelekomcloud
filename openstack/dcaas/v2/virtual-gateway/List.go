package virtual_gateway

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type ListOpts struct {
	Limit        int      `q:"limit"`
	Marker       string   `q:"marker"`
	PageReverse  bool     `q:"page_reverse"`
	ID           []string `q:"id"`
	TenantID     []string `q:"tenant_id"`
	Name         []string `q:"name"`
	Description  []string `q:"description"`
	VPCID        []string `q:"vpc_id"`
	LocalEPGroup []string `q:"local_ep_group_id"`
	DeviceID     []string `q:"device_id"`
	Type         []string `q:"type"`
	Status       []string `q:"status"`
	AdminStateUp []string `q:"admin_state_up"`
}

// List is used to obtain the virtual gateway list
func List(c *golangsdk.ServiceClient, opts ListOpts) ([]VirtualGateway, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v2.0/{project_id}/virtual_gateways
	raw, err := c.Get(c.ServiceURL("dcaas", "virtual-gateways")+q.String(), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []VirtualGateway
	err = extract.IntoSlicePtr(raw.Body, &res, "virtual_gateways")
	return res, err
}

type VirtualGateway struct {
	ID                 string `json:"id"`
	TenantID           string `json:"tenant_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	VPCID              string `json:"vpc_id"`
	LocalEPGroupID     string `json:"local_ep_group_id"`
	DeviceID           string `json:"device_id"`
	RedundantDeviceID  string `json:"redundant_device_id"`
	Type               string `json:"type"`
	IPSecBandwidth     int    `json:"ipsec_bandwidth"`
	Status             string `json:"status"`
	AdminStateUp       bool   `json:"admin_state_up"`
	BGPASN             int    `json:"bgp_asn"`
	RegionID           string `json:"region_id"`
	LocalEPGroupIPv6ID string `json:"local_ep_group_ipv6_id"`
}
