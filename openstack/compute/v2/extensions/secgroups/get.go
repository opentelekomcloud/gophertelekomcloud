package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get will return details for a particular security group.
func Get(client *golangsdk.ServiceClient, id string) (*SecurityGroup, error) {
	raw, err := client.Get(client.ServiceURL("os-security-groups", id), nil, nil)
	return extra(err, raw)
}
