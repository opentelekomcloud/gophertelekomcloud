package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// AddNodesOpts defines options to add master or client nodes .
type AddNodesOpts struct {
	// NodeSize - Number of nodes. The value range is 1 to 32.
	// If the node type is ess-master, the number of nodes must be an odd number in the range 3 to 10.
	// If the node type is ess-client, the number of nodes must be in the range 1 to 32.
	NodeSize int `json:"node_size" required:"true"`
	// Flavor - Flavor ID.
	Flavor string `json:"flavor_ref" required:"true"`
	// Type of the volume.
	// One of:
	//   - `COMMON`: Common I/O
	//   - `HIGH`: High I/O
	//   - `ULTRAHIGH`: Ultra-high I/O
	VolumeType string `json:"volume_type" required:"true"`
}

// AddClusterNodes function lets you add master and client nodes to a cluster.
func AddClusterNodes(client *golangsdk.ServiceClient, clusterID string, NodeType string, opts AddNodesOpts) (*AddNodesResponse, error) {

	b, err := build.RequestBody(opts, "type")
	if err != nil {
		return nil, err
	}

	// POST /v1.0/{project_id}/clusters/{cluster_id}/type/{type}/independent
	url := client.ServiceURL("clusters", clusterID, "type", NodeType, "independent")

	raw, err := client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res AddNodesResponse
	err = extract.IntoStructPtr(raw.Body, &res, "")

	return &res, err
}

type AddNodesResponse struct {
	ID string `json:"id"`
}
