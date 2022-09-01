package nodes

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create accepts a CreateOpts struct and uses the values to create a new
// logical node.
func Create(client *golangsdk.ServiceClient, clusterID string, opts CreateOpts) (r CreateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("clusters", clusterID, "nodes"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	return
}
