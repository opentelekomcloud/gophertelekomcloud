package customer_gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOpts struct {
	// Specifies the name of a customer gateway.
	// If this parameter is not specified, a name in the format of cgw-**** is automatically generated, for example, cgw-21a3.
	// The value is a string of 1 to 64 characters, which can contain digits, letters, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name,omitempty"`
	// Specifies the identifier type of customer gateway.
	// Value range:
	// ip
	// fqdn (currently not supported)
	IdType string `json:"id_type,omitempty"`
	// Specifies the identifier of a customer gateway.
	// The value is a string of 1 to 128 characters. When id_type is set to ip,
	// the value is an IPv4 address in dotted decimal notation, for example, 192.168.45.7.
	// When id_type is set to fqdn, the value is a string of characters that can contain uppercase letters,
	// lowercase letters, digits, and special characters. Spaces and the following special characters are not supported: & < > [ ] \ ?.
	IdValue string `json:"id_value,omitempty"`
	// Specifies the BGP AS number of the customer gateway.
	// The value ranges from 1 to 4294967295.
	// Set this parameter only when id_type is set to ip.
	BgpAsn *int `json:"bgp_asn,omitempty"`
	// Specifies a tag list.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*CustomerGateway, error) {
	b, err := build.RequestBody(opts, "customer_gateway")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("customer-gateways"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202, 201},
	})
	if err != nil {
		return nil, err
	}

	var res CustomerGateway
	return &res, extract.IntoStructPtr(raw.Body, &res, "customer_gateway")
}

type CustomerGateway struct {
	// Specifies a customer gateway ID.
	ID string `json:"id"`
	// Specifies a customer gateway name. If no customer gateway name is specified, the system automatically generates one.
	Name string `json:"name"`
	// Specifies the network type of the VPN gateway.
	IdType string `json:"id_type"`
	// Specifies the association mode.
	IdValue string `json:"id_value"`
	// Specifies the BGP AS number of the customer gateway. This parameter is available only when id_type is set to ip.
	BgpAsn int `json:"bgp_asn"`
	// Specifies a tag list.
	Tags []tags.ResourceTag `json:"tags"`
	// Specifies the time when the customer gateway is created.
	CreatedAt string `json:"created_at"`
	// Specifies the last update time.
	UpdatedAt string `json:"updated_at"`
	// Specifies the routing mode.
	// Value range:
	// static: static routing mode
	// bgp: BGP routing mode
	RouteMode string `json:"route_mode"`
	// Specifies the IP address of the customer gateway.
	Ip string `json:"ip"`
}
