package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateIpList is used to create or update the ip list of specific ip group.
func UpdateIpList(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (*IpGroup, error) {
	b, err := build.RequestBody(opts, "ipgroup")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/elb/ipgroups/{ipgroup_id}/iplist/create-or-update
	raw, err := c.Post(c.ServiceURL("ipgroups", id, "iplist", "create-or-update"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
