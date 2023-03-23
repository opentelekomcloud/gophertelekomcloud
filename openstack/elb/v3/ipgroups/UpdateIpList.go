package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateIpList is used to create or update the ip list of specific ip group.
func UpdateIpList(c *golangsdk.ServiceClient, id string, opts UpdateOpts) (*IpGroup, error) {
	b, err := build.RequestBody(opts, "ipgroup")
	if err != nil {
		return nil, err
	}
	url := c.ServiceURL("ipgroups", id, "iplist", "create-or-update")
	raw, err := c.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res IpGroup
	err = extract.IntoStructPtr(raw.Body, &res, "ipgroup")
	return &res, err
}
