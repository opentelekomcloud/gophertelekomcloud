package nodepools

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

// List returns collection of node pools.
func List(client *golangsdk.ServiceClient, clusterID string, opts ListOpts) ([]NodePool, error) {
	// GET /api/v3/projects/{project_id}/clusters/{cluster_id}/nodepools
	raw, err := client.Get(client.ServiceURL("clusters", clusterID, "nodepools"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return nil, err
	}

	var res ListNodePool
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return FilterNodePools(res.NodePools, opts), nil
}

// ListNodePool - Describes the Node Pool Structure of cluster
type ListNodePool struct {
	// API type, fixed value "List"
	Kind string `json:"kind"`
	// API version, fixed value "v3"
	Apiversion string `json:"apiVersion"`
	// all Node Pools
	NodePools []NodePool `json:"items"`
}

// FilterNodes filters a list of Nodes based on the given ListOpts.
func FilterNodePools(nodepools []NodePool, opts ListOpts) []NodePool {
	var filteredNodepools []NodePool

	for _, nodepool := range nodepools {
		if matchesFilters(nodepool, opts) {
			filteredNodepools = append(filteredNodepools, nodepool)
		}
	}
	return filteredNodepools
}

// matchesFilters checks if a node satisfies the filtering criteria in ListOpts.
func matchesFilters(nodepool NodePool, opts ListOpts) bool {
	// Check each filter explicitly
	if opts.Name != "" && nodepool.Metadata.Name != opts.Name {
		return false
	}
	if opts.Uid != "" && nodepool.Metadata.Id != opts.Uid {
		return false
	}
	if opts.Phase != "" && nodepool.Status.Phase != opts.Phase {
		return false
	}

	return true
}
