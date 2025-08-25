package natgateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateGatewayOpts struct {
	// Specifies the private NAT gateway name.
	// Only digits, letters, underscores (_), and hyphens (-) are allowed.
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the private NAT gateway.
	// The description can contain up to 255 characters and cannot contain angle brackets (<>).
	Description string `json:"description,omitempty"`
	// Specifies the private NAT gateway specifications.
	// The value can be: Small, Medium, Large, Extra-large.
	// Default value: Small.
	Spec string `json:"spec,omitempty"`
}

// This function is used to create a Private NAT gateway.
func Update(client *golangsdk.ServiceClient, gatewayId string, opts UpdateGatewayOpts) (*GatewayCommonResponse, error) {
	b, err := build.RequestBody(opts, "gateway")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/private-nat/gateways/{gateway_id}
	raw, err := client.Put(client.ServiceURL("private-nat", "gateways", gatewayId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res GatewayCommonResponse
	return &res, extract.Into(raw.Body, &res)
}
