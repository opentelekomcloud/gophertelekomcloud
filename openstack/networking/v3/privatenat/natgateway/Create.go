package natgateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateGatewayOpts struct {
	// Specifies the private NAT gateway name.
	// Only digits, letters, underscores (_), and hyphens (-) are allowed.
	Name string `json:"name" required:"true"`
	// Provides supplementary information about the private NAT gateway.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`
	// Specifies the private NAT gateway specifications.
	// The value can be: Small, Medium, Large, Extra-large.
	// Default value: Small.
	Spec string `json:"spec,omitempty"`
	// Specifies the VPC where the private NAT gateway works.
	DownlinkVpcs []DownlinkVpcOption `json:"downlink_vpcs" required:"true"`
	// Specifies the tag list.
	Tags []TagOptions `json:"tags,omitempty"`
	// Specifies the ID of the enterprise project associated with the private NAT gateway.
	// Default value: 0.
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`
}

type DownlinkVpcOption struct {
	// Specifies the ID of the subnet where the private NAT gateway works.
	VirSubnetID string `json:"virsubnet_id" required:"true"`
	// Specifies the private IP address of the private NAT gateway.
	NgPortIPAddress string `json:"ngport_ip_address,omitempty"`
}

type TagOptions struct {
	// Specifies the tag key. A key can contain up to 128 Unicode characters.
	// Key cannot be left blank.
	Key string `json:"key" required:"true"`
	// Specifies the tag value. Each value can contain up to 255 Unicode characters.
	Value string `json:"value" required:"true"`
}

// This function is used to create a Private NAT gateway.
func Create(client *golangsdk.ServiceClient, opts CreateGatewayOpts) (*GatewayCommonResponse, error) {
	b, err := build.RequestBody(opts, "gateway")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/private-nat/gateways
	raw, err := client.Post(client.ServiceURL("private-nat", "gateways"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res GatewayCommonResponse
	return &res, extract.Into(raw.Body, &res)
}
