package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts represents options for updating a IpGroup.
type UpdateOpts struct {
	// Specifies the IP address group name.
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the IP address group.
	Description string `json:"description,omitempty"`
	// Lists the IP addresses in the IP address group.
	IpList []IpGroupOption `json:"ip_list,omitempty"`
}

// Update is an operation which modifies the attributes of the specified IpGroup.
func Update(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (*IpGroup, error) {
	b, err := build.RequestBody(opts, "ipgroup")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/elb/ipgroups/{ipgroup_id}
	raw, err := c.Put(c.ServiceURL("ipgroups", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})

	return extra(err, raw)
}
