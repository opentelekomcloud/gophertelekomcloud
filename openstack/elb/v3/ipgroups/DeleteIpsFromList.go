package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// DeleteIpFromList is used to create or update the ip list of specific ip group.
func DeleteIpFromList(c *golangsdk.ServiceClient, id string, opts BatchDeleteOpts) (*IpGroup, error) {
	b, err := build.RequestBody(opts, "ipgroup")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/elb/ipgroups/{ipgroup_id}/iplist/batch-delete
	raw, err := c.Post(c.ServiceURL("ipgroups", id, "iplist", "batch-delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}

// BatchDeleteOpts contains all the values needed to perform BatchDelete on the IP address group.
type BatchDeleteOpts struct {
	// Specifies IP addresses that will be deleted from an IP address group in batches.
	IpList []IpList `json:"ip_list,omitempty"`
}
type IpList struct {
	// Specifies the IP addresses in the IP address group.
	Ip string `json:"ip" required:"true"`
}
