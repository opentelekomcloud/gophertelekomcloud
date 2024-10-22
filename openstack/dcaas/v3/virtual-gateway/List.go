package virtual_gateway

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ListOpts struct {
	// virtual gateway by ID
	ID string `q:"id,omitempty"`
	// Specifies the number of records returned on each page. Value range: 1-2000
	Limit int `q:"limit,omitempty"`
	// Specifies the ID of the last resource record on the previous page. If this parameter is left blank, the first page is queried.
	// This parameter must be used together with limit.
	Marker string `q:"marker,omitempty"`
	// Specifies the list of fields to be displayed.
	Fields []interface{} `q:"fields,omitempty"`
	// Specifies the sorting order of returned results. The value can be asc (default) or desc.
	SortDir string `q:"sort_dir,omitempty"`
	// Specifies the field for sorting.
	SortKey string `q:"sort_key,omitempty"`
	// Specifies the VPC ID by which virtual gateways are queried.
	VpcId string `q:"vpc_id,omitempty"`
}

// List is used to obtain the virtual gateway list
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]VirtualGateway, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("dcaas", "virtual-gateways").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/virtual-gateways
	raw, err := client.Get(client.ServiceURL(url.String()), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []VirtualGateway
	err = extract.IntoSlicePtr(raw.Body, &res, "virtual_gateways")
	return res, err
}

type VirtualGateway struct {
	// The ID of the virtual gateway.
	ID string `json:"id"`
	// The ID of the VPC connected to the virtual gateway.
	VpcId string `json:"vpc_id"`
	// The project ID to which the virtual gateway belongs.
	TenantId string `json:"tenant_id"`
	// Specifies the name of the virtual gateway.
	// The valid length is limited from 0 to 64, only chinese and english letters, digits, hyphens (-), underscores (_)
	// and dots (.) are allowed.
	// The name must start with a chinese or english letter, and the Chinese characters must be in **UTF-8** or
	// **Unicode** format.
	Name string `json:"name"`
	// Specifies the description of the virtual gateway.
	// The description contain a maximum of 64 characters and the angle brackets (< and >) are not allowed.
	// Chinese characters must be in **UTF-8** or **Unicode** format.
	Description string `json:"description"`
	// The type of virtual gateway.
	Type string `json:"type"`
	// The list of IPv4 subnets from the virtual gateway to access cloud services, which is usually the CIDR block of
	// the VPC.
	LocalEpGroup []string `json:"local_ep_group"`
	// The list of IPv6 subnets from the virtual gateway to access cloud services, which is usually the CIDR block of
	// the VPC.
	LocalEpGroupIpv6 []string `json:"local_ep_group_ipv6"`
	// The current status of the virtual gateway.
	Status string `json:"status"`
	// The local BGP ASN of the virtual gateway.
	BgpAsn int `json:"bgp_asn"`
	// The key/value pairs to associate with the virtual gateway.
	Tags []tags.ResourceTag `json:"tags"`
}
