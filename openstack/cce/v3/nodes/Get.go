package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves a particular nodes based on its unique ID and cluster ID.
func Get(client *golangsdk.ServiceClient, clusterID, nodeID string) (*Nodes, error) {
	// GET /api/v3/projects/{project_id}/clusters/{cluster_id}/nodes/{node_id}
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "nodes", nodeID), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return nil, err
	}

	var res Nodes
	return &res, extract.Into(raw.Body, &res)
}
