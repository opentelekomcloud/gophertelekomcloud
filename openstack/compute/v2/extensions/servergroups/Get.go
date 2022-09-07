package servergroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns data about a previously created ServerGroup.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(getURL(client, id), nil, nil)
	return
}
