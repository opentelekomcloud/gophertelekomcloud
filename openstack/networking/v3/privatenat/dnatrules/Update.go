package dnatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePrivateDnatOpts struct {
	// Provides supplementary information about the DNAT rule.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`
	// Specifies the ID of the transit IP address.
	TransitIpId string `json:"transit_ip_id,omitempty"`
	// Specifies the port ID of the resource that the NAT gateway is bound to.
	// The resource can be a compute instance, load balancer (v2 or v3), or virtual IP address.
	// Either this parameter or private_ip_address must be specified.
	NetworkInterfaceId string `json:"network_interface_id,omitempty"`
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

// This function is used to update a DNAT rule.
func Update(client *golangsdk.ServiceClient, ruleId string, opts UpdatePrivateDnatOpts) (*DnatCommonResponse, error) {
	b, err := build.RequestBody(opts, "dnat_rule")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/private-nat/dnat-rules/{dnat_rule_id}
	raw, err := client.Put(client.ServiceURL("private-nat", "dnat-rules", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res DnatCommonResponse
	return &res, extract.Into(raw.Body, &res)
}
