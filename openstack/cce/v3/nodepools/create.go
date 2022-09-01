package nodepools

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create accepts a CreateOpts struct and uses the values to create a new
// logical node pool.
func Create(client *golangsdk.ServiceClient, clusterid string, opts CreateOpts) (r CreateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("clusters", clusterid, "nodepools"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	return
}
