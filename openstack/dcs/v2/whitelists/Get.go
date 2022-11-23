package whitelists

import (
	"strings"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get the instance whitelist groups by instance id
func Get(client *golangsdk.ServiceClient, id string) (*Whitelist, error) {
	url := client.ServiceURL("instance", id, "whitelist")
	raw, err := client.Get(strings.Replace(url, "v1.0", "v2", 1), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Whitelist
	err = extract.Into(raw.Body, &res)
	return &res, err
}

// Whitelist is a struct that contains all the whitelist parameters.
type Whitelist struct {
	// instance id
	InstanceID string `json:"instance_id"`
	// enable or disable the whitelists
	Enable bool `json:"enable_whitelist"`
	// list of whitelist groups
	Groups []WhitelistGroup `json:"whitelist"`
}

// WhitelistGroup is a struct that contains the whitelist parameters.
type WhitelistGroup struct {
	// the group name
	GroupName string `json:"group_name"`
	// list of IP address or range
	IPList []string `json:"ip_list"`
}
