package snatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]SnatRule, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("snat_rules").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2.0/snat_rules
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []SnatRule
	err = extract.IntoSlicePtr(raw.Body, &res, "snat_rules")
	return res, err
}

type ListOpts struct {
	// Specifies the SNAT rule ID.
	Id string `q:"id,omitempty"`
	// Specifies the number of records returned on each page.
	Limit int `q:"limit,omitempty"`
	// Specifies the project ID.
	ProjectId string `q:"tenant_id,omitempty"`
	// Specifies the NAT gateway ID.
	NatGatewayId string `q:"nat_gateway_id,omitempty"`
	// Specifies the network ID used by the SNAT rule.
	NetworkId string `q:"network_id,omitempty"`
	// Specifies a subset of the VPC subnet CIDR block or a CIDR block of Direct Connect connection.
	Cidr string `q:"cidr,omitempty"`
	// 0: Either network_id or cidr can be specified in a VPC.
	// 1: Only cidr can be specified over a Direct Connect connection.
	// If no value is entered, the default value 0 (VPC) is used.
	SourceType string `q:"source_type,omitempty"`
	// Specifies the EIP ID.
	FloatingIpId string `q:"floating_ip_id,omitempty"`
	// Specifies the EIP address.
	FloatingIpAddress string `q:"floating_ip_address,omitempty"`
	// Provides supplementary information about the SNAT rule.
	Description string `q:"description,omitempty"`
	// Specifies the status of the SNAT rule.
	Status string `q:"status,omitempty"`
	// Specifies whether the SNAT rule is enabled or disabled.
	// The value can be:
	// true: The SNAT rule is enabled.
	// false: The SNAT rule is disabled.
	AdminStateUp bool `q:"admin_state_up,omitempty"`
	// Specifies when the SNAT rule was created (UTC time).
	// Its value rounds to 6 decimal places for seconds. The format is yyyy-mm-dd hh:mm:ss.
	CreatedAt bool `q:"created_at,omitempty"`
}
