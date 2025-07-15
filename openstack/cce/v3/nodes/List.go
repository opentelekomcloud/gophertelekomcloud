package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts allows the filtering of list data using given parameters.
type ListOpts struct {
	Name  string `json:"name"`
	Uid   string `json:"uid"`
	Phase string `json:"phase"`
}

// List returns collection of nodes.
func List(client *golangsdk.ServiceClient, clusterID string, opts ListOpts) ([]Nodes, error) {
	// GET /api/v3/projects/{project_id}/clusters/{cluster_id}/nodes
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "nodes"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return nil, err
	}

	var res ListNode
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return FilterNodes(res.Nodes, opts), nil
}

// ListNode describes the Node Structure of cluster
type ListNode struct {
	// API type, fixed value "List"
	Kind string `json:"kind"`
	// API version, fixed value "v3"
	Apiversion string `json:"apiVersion"`
	// all Clusters
	Nodes []Nodes `json:"items"`
}

// FilterNodes filters a list of Nodes based on the given ListOpts.
func FilterNodes(nodes []Nodes, opts ListOpts) []Nodes {
	var filteredNodes []Nodes

	for _, node := range nodes {
		if matchesFilters(node, opts) {
			filteredNodes = append(filteredNodes, node)
		}
	}
	return filteredNodes
}

// matchesFilters checks if a node satisfies the filtering criteria in ListOpts.
func matchesFilters(node Nodes, opts ListOpts) bool {
	// Check each filter explicitly
	if opts.Name != "" && node.Metadata.Name != opts.Name {
		return false
	}
	if opts.Uid != "" && node.Metadata.Id != opts.Uid {
		return false
	}
	if opts.Phase != "" && node.Status.Phase != opts.Phase {
		return false
	}

	return true
}
