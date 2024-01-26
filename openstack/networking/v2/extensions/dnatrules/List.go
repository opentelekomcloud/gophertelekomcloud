package dnatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]DnatRule, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v2.0/dnat_rules
	url := client.ServiceURL("dnat_rules") + query.String()
	raw, err := client.Get(url, nil, nil)
	if err != nil {
		return nil, err
	}

	var res []DnatRule
	err = extract.IntoSlicePtr(raw.Body, &res, "dnat_rules")
	return res, err
}

type ListOpts struct {
	// Specifies the DNAT rule ID.
	Id string `q:"id,omitempty"`
	// Specifies the number of records returned on each page.
	Limit int `q:"limit,omitempty"`
	// Specifies the project ID.
	ProjectId string `q:"tenant_id,omitempty"`
	// Specifies the NAT gateway ID.
	NatGatewayId string `q:"nat_gateway_id,omitempty"`
	// Specifies the port ID of the cloud server (ECS or BMS).
	PortId string `q:"port_id,omitempty"`
	// Private IP address
	PrivateIp string `q:"private_ip,omitempty"`
	// Specifies the port number used by the cloud server (ECS or BMS) to provide services for external systems.
	InternalServicePort int `q:"internal_service_port,omitempty"`
	// Specifies the EIP ID.
	FloatingIpId string `q:"floating_ip_id,omitempty"`
	// Specifies the EIP address.
	FloatingIpAddress string `q:"floating_ip_address,omitempty"`
	// Specifies the port for providing services for external systems.
	ExternalServicePort int `q:"external_service_port,omitempty"`
	// Specifies the protocol. TCP, UDP, and ANY are supported.
	// The protocol number of TCP, UDP, and ANY are 6, 17, and 0, respectively.
	Protocol string `q:"protocol,omitempty"`
	// Provides supplementary information about the DNAT rule.
	Description string `q:"description,omitempty"`
	// Specifies the status of the DNAT rule.
	Status string `q:"status,omitempty"`
	// Specifies whether the DNAT rule is enabled or disabled.
	// The value can be:
	// true: The DNAT rule is enabled.
	// false: The DNAT rule is disabled.
	AdminStateUp bool `q:"admin_state_up,omitempty"`
	// Specifies when the DNAT rule was created (UTC time).
	// Its value rounds to 6 decimal places for seconds. The format is yyyy-mm-dd hh:mm:ss.
	CreatedAt bool `q:"created_at,omitempty"`
}
