package ipgroups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DeleteIpFromList is used to create or update the ip list of specific ip group.
func DeleteIpFromList(c *golangsdk.ServiceClient, id string, opts BatchDeleteOpts) (*IpGroup, error) {
	b, err := build.RequestBody(opts, "ipgroup")
	if err != nil {
		return nil, err
	}
	url := c.ServiceURL("ipgroups", id, "iplist", "batch-delete")
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

// BatchDeleteOpts contains all the values needed to perform BatchDelete on the IP address group.
type BatchDeleteOpts struct {
	// Specifies IP addresses that will be deleted from an IP address group in batches.
	IpList []IpList `json:"ip_list,omitempty"`
}
type IpList struct {
	// Specifies the IP addresses in the IP address group.
	Ip string `json:"ip" required:"true"`
}
