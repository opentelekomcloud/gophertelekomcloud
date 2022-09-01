package nodes

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves a particular nodes based on its unique ID and cluster ID.
func Get(client *golangsdk.ServiceClient, clusterID, nodeID string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "nodes", nodeID), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts.MoreHeaders, JSONBody: nil,
	})
	return
}
