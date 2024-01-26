package dnatrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains all the values needed to create a new dnat rule
// resource.
type CreateOpts struct {
	NatGatewayID        string `json:"nat_gateway_id" required:"true"`
	PortID              string `json:"port_id,omitempty"`
	PrivateIp           string `json:"private_ip,omitempty"`
	InternalServicePort *int   `json:"internal_service_port" required:"true"`
	FloatingIpID        string `json:"floating_ip_id" required:"true"`
	ExternalServicePort *int   `json:"external_service_port" required:"true"`
	Protocol            string `json:"protocol" required:"true"`
}

// Create will create a new Waf Certificate on the values in CreateOpts.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*DnatRule, error) {
	b, err := build.RequestBody(opts, "dnat_rule")
	if err != nil {
		return nil, err
	}

	// POST /v2.0/dnat_rules
	raw, err := client.Post(client.ServiceURL("dnat_rules"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201},
		})
	if err != nil {
		return nil, err
	}

	var res DnatRule
	err = extract.IntoStructPtr(raw.Body, &res, "dnat_rule")
	return &res, err
}

type DnatRule struct {
	// Specifies the DNAT rule ID.
	ID string `json:"id"`
	// Specifies the project ID.
	ProjectId string `json:"tenant_id"`
	// Specifies the NAT gateway ID.
	NatGatewayId string `json:"nat_gateway_id"`
	// Specifies the port ID of the cloud server (ECS or BMS).
	// This parameter is used in the VPC scenario, where this parameter or private_ip must be specified.
	PortId string `json:"port_id"`
	// Specifies the IP address of an on-premises network connected by a Direct Connect connection.
	// This parameter is used in the Direct Connect scenario. This parameter and port_id are alternative.
	PrivateIp string `json:"private_ip"`
	// Specifies the port number used by the cloud server (ECS or BMS) to provide services for external systems.
	InternalServicePort int `json:"internal_service_port"`
	// Specifies the EIP ID.
	FloatingIpId string `json:"floating_ip_id"`
	// Specifies the EIP address.
	FloatingIpAddress string `json:"floating_ip_address"`
	// Specifies the port for providing services for external systems.
	ExternalServicePort int `json:"external_service_port"`
	// Specifies the protocol. TCP, UDP, and ANY are supported.
	// The protocol number of TCP, UDP, and ANY are 6, 17, and 0, respectively.
	Protocol string `json:"protocol"`
	// Provides supplementary information about the DNAT rule.
	Description string `json:"description"`
	// Specifies the status of the DNAT rule.
	Status string `json:"status"`
	// Specifies whether the NAT gateway is up or down.
	// The value can be:
	// true: The DNAT rule is enabled.
	// false: The DNAT rule is disabled.
	AdminStateUp *bool `json:"admin_state_up"`
	// Specifies when the DNAT rule was created (UTC time).
	// Its value rounds to 6 decimal places for seconds. The format is yyyy-mm-dd hh:mm:ss.
	CreatedAt string `json:"created_at"`
}
