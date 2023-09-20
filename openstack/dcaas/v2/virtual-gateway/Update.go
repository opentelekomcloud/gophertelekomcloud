package virtual_gateway

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts represents options for updating a VirtualGateway.
type UpdateOpts struct {
	// Provides supplementary information about the virtual gateway.
	Description string `json:"description,omitempty"`
	// Specifies the virtual gateway name.
	Name string `json:"name,omitempty"`
	// Specifies the ID of the local endpoint group that records CIDR blocks of the VPC subnets.
	LocalEndpointGroupId string `json:"local_ep_group_id,omitempty"`
}

// Update is an operation which modifies the attributes of the specified
// VirtualGateway.
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "virtual_gateway")
	if err != nil {
		return
	}

	_, err = c.Put(c.ServiceURL("dcaas", "virtual-gateways", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})

	return
}
