package clusters

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create accepts a CreateOpts struct and uses the values to create a new
// logical cluster.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (r CreateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("clusters"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	return
}
