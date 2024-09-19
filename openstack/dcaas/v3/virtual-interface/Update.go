package virtual_interface

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts represents options for updating a Virtual Interface.
type UpdateOpts struct {
	// Provides supplementary information about the virtual interface.
	Description string `json:"description,omitempty"`
	// Specifies the virtual interface name.
	Name string `json:"name,omitempty"`
	// Specifies the virtual interface bandwidth.
	Bandwidth int `json:"bandwidth,omitempty"`
	// Specifies the ID of the remote endpoint group that records the CIDR blocks used by the on-premises network.
	RemoteEndpointGroupId string `json:"remote_ep_group_id,omitempty"`
}

// Update is an operation which modifies the attributes of the specified
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "virtual_interface")
	if err != nil {
		return
	}

	_, err = c.Put(c.ServiceURL("dcaas", "virtual-interfaces", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})

	return
}
