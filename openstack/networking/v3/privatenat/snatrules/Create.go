package snatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreatePrivateSnatOpts struct {
	// Specifies the private NAT gateway ID.
	GatewayId string `json:"gateway_id" required:"true"`
	// Specifies the CIDR block that matches the SNAT rule.
	// Constraint: Either this parameter or virsubnet_id must be specified.
	Cidr string `json:"cidr,omitempty"`
	// Specifies the ID of the subnet that matches the SNAT rule.
	// Constraint: Either this parameter or cidr must be specified.
	VirSubnetId string `json:"virsubnet_id,omitempty"`
	// Provides supplementary information about the SNAT rule.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`
	// Specifies the IDs of the transit IP addresses.
	// Constraints: A maximum number of 20 IDs is allowed.
	TransitIpIds []string `json:"transit_ip_ids" required:"true"`
}

// This function is used to create an SNAT rule.
func Create(client *golangsdk.ServiceClient, opts CreatePrivateSnatOpts) (*SnatCommonResponse, error) {
	b, err := build.RequestBody(opts, "snat_rule")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/private-nat/snat-rules
	raw, err := client.Post(client.ServiceURL("private-nat", "snat-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res SnatCommonResponse
	return &res, extract.Into(raw.Body, &res)
}
