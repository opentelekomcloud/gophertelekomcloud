package nodes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
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
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("clusters", clusterID, "nodes", nodeID), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Nodes
	err = extract.Into(raw.Body, &res)
	return &res, err
}
