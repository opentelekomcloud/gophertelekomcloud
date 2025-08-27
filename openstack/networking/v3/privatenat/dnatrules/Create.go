package dnatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreatePrivateDnatOpts struct {
	// Provides supplementary information about the DNAT rule.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`
	// Specifies the ID of the transit IP address.
	TransitIpId string `json:"transit_ip_id" required:"true"`
	// Specifies the port ID of the resource that the NAT gateway is bound to.
	// The resource can be a compute instance, load balancer (v2 or v3), or virtual IP address.
	// Either this parameter or private_ip_address must be specified.
	NetworkInterfaceId string `json:"network_interface_id,omitempty"`
	// Specifies the private NAT gateway ID.
	GatewayId string `json:"gateway_id" required:"true"`
	// Specifies the port IP address that the NAT gateway uses.
	// The resource can be a compute instance, load balancer (v2 or v3), or virtual IP address.
	// Either this parameter or network_interface_id must be specified.
	PrivateIpAddress string `json:"private_ip_address,omitempty"`
	// Specifies the protocol type. TCP, UDP, and ANY are supported.
	// Enumeration values: tcp, udp, any
	Protocol string `json:"protocol,omitempty"`
	// Specifies the port number of the resource, which can be a compute instance,
	// load balancer (v2 or v3), or virtual IP address.
	// Value range: 0-65535, Default: 0
	InternalServicePort string `json:"internal_service_port,omitempty"`
	// Specifies the port number of the transit IP address.
	// Value range: 0-65535, Default: 0
	TransitServicePort string `json:"transit_service_port,omitempty"`
}

// This function is used to create a DNAT rule.
func Create(client *golangsdk.ServiceClient, opts CreatePrivateDnatOpts) (*DnatCommonResponse, error) {
	b, err := build.RequestBody(opts, "dnat_rule")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/private-nat/dnat-rules
	raw, err := client.Post(client.ServiceURL("private-nat", "dnat-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res DnatCommonResponse
	return &res, extract.Into(raw.Body, &res)
}
