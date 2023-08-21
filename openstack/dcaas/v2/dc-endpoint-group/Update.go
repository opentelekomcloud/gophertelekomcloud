package dc_endpoint_group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Provides supplementary information about the Direct Connect endpoint group.
	Description string `json:"description,omitempty"`
	// Specifies the name of the Direct Connect endpoint group.
	Name string `json:"name,omitempty"`
}

// Update is an operation which modifies the attributes of the specified Direct Connect endpoint group
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "dc_endpoint_group")
	if err != nil {
		return
	}

	_, err = c.Put(c.ServiceURL("dcaas", "dc-endpoint-groups", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})

	return
}
