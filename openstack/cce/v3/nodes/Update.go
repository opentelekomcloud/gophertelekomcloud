package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts contains all the values needed to update a new node
type UpdateOpts struct {
	Metadata UpdateMetadata `json:"metadata,omitempty"`
}

type UpdateMetadata struct {
	Name string `json:"name,omitempty"`
}

// Update allows nodes to be updated.
func Update(client *golangsdk.ServiceClient, clusterID, nodeID string, opts UpdateOpts) (*Nodes, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /api/v3/projects/{project_id}/clusters/{cluster_id}/nodes/{node_id}
	raw, err := client.Put(client.ServiceURL("clusters", clusterID, "nodes", nodeID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Nodes
	return &res, extract.Into(raw.Body, &res)
}
