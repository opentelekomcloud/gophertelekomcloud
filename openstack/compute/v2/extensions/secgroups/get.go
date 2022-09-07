package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get will return details for a particular security group.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(resourceURL(client, id), nil, nil)
	return
}
