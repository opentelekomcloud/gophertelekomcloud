package nodepools

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves a particular node pool based on its unique ID and cluster ID.
func Get(client *golangsdk.ServiceClient, clusterid, nodepoolid string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterid, "nodepools", nodepoolid), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
