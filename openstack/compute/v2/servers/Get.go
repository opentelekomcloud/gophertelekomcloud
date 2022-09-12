package servers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get requests details on a single server, by ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("servers", id), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 203},
	})
	return
}
