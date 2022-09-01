package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves a particular cluster based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("clusters", id), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts, JSONBody: nil,
	})
	return
}
