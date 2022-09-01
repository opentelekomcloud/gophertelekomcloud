package nodes

import "github.com/opentelekomcloud/gophertelekomcloud"

// Update allows nodes to be updated.
func Update(client *golangsdk.ServiceClient, clusterID, nodeID string, opts UpdateOpts) (r UpdateResult) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("clusters", clusterID, "nodes", nodeID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
