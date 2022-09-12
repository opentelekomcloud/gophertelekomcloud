package servergroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns data about a previously created ServerGroup.
func Get(client *golangsdk.ServiceClient, id string) (*ServerGroup, error) {
	raw, err := client.Get(client.ServiceURL("os-server-groups", id), nil, nil)
	return extra(err, raw)
}
