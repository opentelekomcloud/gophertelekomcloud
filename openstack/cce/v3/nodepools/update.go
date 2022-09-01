package nodepools

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update allows node pools to be updated.
func Update(client *golangsdk.ServiceClient, clusterid, nodepoolid string, opts UpdateOpts) (r UpdateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("clusters", clusterid, "nodepools", nodepoolid), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
