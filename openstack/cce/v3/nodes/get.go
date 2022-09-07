package nodes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

var RequestOpts = map[string]string{"Content-Type": "application/json"}

// Get retrieves a particular nodes based on its unique ID and cluster ID.
func Get(client *golangsdk.ServiceClient, clusterID, nodeID string) (*Nodes, error) {
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "nodes", nodeID), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts, JSONBody: nil,
	})
	if err != nil {
		return nil, err
	}

	var res Nodes
	err = extract.Into(raw.Body, &res)
	return &res, err
}
