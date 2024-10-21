package virtual_interface

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Specifies the name of the virtual interface.
	// The valid length is limited from 0 to 64, only chinese and english letters, digits, hyphens (-), underscores (_)
	// and dots (.) are allowed.
	// The name must start with a chinese or english letter, and the Chinese characters must be in **UTF-8** or
	// **Unicode** format.
	Name string `json:"name,omitempty"`
	// Specifies the description of the virtual interface.
	// The description contain a maximum of 128 characters and the angle brackets (< and >) are not allowed.
	// Chinese characters must be in **UTF-8** or **Unicode** format.
	Description *string `json:"description,omitempty"`
	// The ingress bandwidth size of the virtual interface.
	Bandwidth int `json:"bandwidth,omitempty"`
	// The CIDR list of remote subnets.
	RemoteEpGroup []string `json:"remote_ep_group,omitempty"`
	// The CIDR list of subnets in service side.
	ServiceEpGroup []string `json:"service_ep_group,omitempty"`
	// Whether to enable the Bidirectional Forwarding Detection (BFD) function.
	EnableBfd *bool `json:"enable_bfd,omitempty"`
	// Whether to enable the Network Quality Analysis (NQA) function.
	EnableNqa *bool `json:"enable_nqa,omitempty"`
	// The status of the virtual interface to be changed.
	Status string `json:"status,omitempty"`
}

// Update is an operation which modifies the attributes of the specified
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (*VirtualInterface, error) {
	b, err := build.RequestBody(opts, "virtual_interface")
	if err != nil {
		return nil, err
	}

	raw, err := c.Put(c.ServiceURL("dcaas", "virtual-interfaces", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}
	var res VirtualInterface
	err = extract.IntoStructPtr(raw.Body, &res, "virtual_interface")
	return &res, err
}
