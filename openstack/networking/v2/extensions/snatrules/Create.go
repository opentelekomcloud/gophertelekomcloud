package snatrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the NAT gateway ID.
	NatGatewayID string `json:"nat_gateway_id" required:"true"`
	// Specifies the network ID used by the SNAT rule. This parameter and cidr are alternative.
	NetworkID string `json:"network_id,omitempty"`
	// Specifies the EIP ID. Use commas (,) to separate multiple IDs.
	// The maximum length of the ID is 4,096 bytes.
	// Constraints: The number of EIP IDs cannot exceed 20.
	FloatingIPID string `json:"floating_ip_id" required:"true"`
	// Specifies CIDR, which can be in the format of a network segment or a host IP address.
	// This parameter and network_id are alternative.
	// If source_type is set to 0, cidr must be a subset of the VPC subnet.
	// If source_type is set to 1, cidr must be a CIDR block of your on-premises network connected to the VPC through Direct Connect or Cloud Connect.
	Cidr string `json:"cidr,omitempty"`
	// 0: Either network_id or cidr can be specified in a VPC.
	// 1: Only cidr can be specified over a Direct Connect connection.
	// If no value is entered, the default value 0 (VPC) is used.
	SourceType int `json:"source_type,omitempty"`
	// Provides supplementary information about the SNAT rule.
	Description string `json:"description,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*SnatRule, error) {
	b, err := build.RequestBody(opts, "snat_rule")
	if err != nil {
		return nil, err
	}

	// POST /v2.0/snat_rules
	raw, err := client.Post(client.ServiceURL("snat_rules"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes: []int{201},
		})
	if err != nil {
		return nil, err
	}

	var res SnatRule
	err = extract.IntoStructPtr(raw.Body, &res, "snat_rule")
	return &res, err
}

type SnatRule struct {
	// Specifies the SNAT rule ID.
	ID string `json:"id"`
	// Specifies the NAT gateway ID.
	NatGatewayID string `json:"nat_gateway_id"`
	// Specifies the network ID used by the SNAT rule.
	NetworkID string `json:"network_id"`
	// Specifies the project ID.
	TenantID string `json:"tenant_id"`
	// Specifies the EIP ID. Use commas (,) to separate multiple IDs.
	FloatingIPID string `json:"floating_ip_id"`
	// Specifies the EIP. Use commas (,) to separate multiple EIPs.
	FloatingIPAddress string `json:"floating_ip_address"`
	// Specifies the status of the SNAT rule.
	Status string `json:"status"`
	// Specifies the unfrozen or frozen state.
	// Specifies whether the SNAT rule is enabled or disabled.
	// The state can be:
	// true: The SNAT rule is enabled.
	// false: The SNAT rule is disabled.
	AdminStateUp bool `json:"admin_state_up"`
	// Specifies a subset of the VPC subnet CIDR block or a CIDR block of Direct Connect connection.
	Cidr string `json:"cidr"`
	// 0: Either network_id or cidr can be specified in a VPC.
	// 1: Only cidr can be specified over a Direct Connect connection.
	// If no value is entered, the default value 0 (VPC) is used.
	SourceType interface{} `json:"source_type"`
	// Specifies when the SNAT rule is created (UTC time).
	// Its value rounds to 6 decimal places for seconds.
	// The format is yyyy-mm-dd hh:mm:ss.
	CreatedAt string `json:"created_at"`
}
