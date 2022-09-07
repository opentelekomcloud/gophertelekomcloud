package defsecrules

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get will return details for a particular default rule.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(resourceURL(client, id), nil, nil)
	return
}
