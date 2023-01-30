package whitelists

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// WhitelistOpts is a struct that contains all the parameters.
type WhitelistOpts struct {
	// enable or disable the whitelists
	Enable *bool `json:"enable_whitelist" required:"true"`
	// list of whitelist groups
	Groups []WhitelistGroupOpts `json:"whitelist" required:"true"`
}

// WhitelistGroupOpts is a struct that contains all the whitelist parameters.
type WhitelistGroupOpts struct {
	// the group name
	GroupName string `json:"group_name" required:"true"`
	// list of IP address or range
	IPList []string `json:"ip_list" required:"true"`
}

// Put an instance whitelist with given parameters.
func Put(client *golangsdk.ServiceClient, id string, opts WhitelistOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	url := client.ServiceURL("instance", id, "whitelist")
	_, err = client.Put(strings.Replace(url, "v1.0", "v2", 1), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
