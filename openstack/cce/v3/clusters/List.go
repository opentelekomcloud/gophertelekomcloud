package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts allows the filtering of list data using given parameters.
type ListOpts struct {
	Name  string `json:"name"`
	ID    string `json:"uuid"`
	Type  string `json:"type"`
	VpcID string `json:"vpc"`
	Phase string `json:"phase"`
}

// List returns collection of clusters.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Clusters, error) {
	// GET /api/v3/projects/{project_id}/clusters
	raw, err := client.Get(client.ServiceURL("clusters"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
		JSONBody:    nil,
	})
	if err != nil {
		return nil, err
	}

	var res ListCluster
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return FilterClusters(res.Clusters, opts), nil
}

type ListCluster struct {
	// API type, fixed value Cluster
	Kind string `json:"kind"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion"`
	// all Clusters
	Clusters []Clusters `json:"items"`
}

func FilterClusters(clusters []Clusters, opts ListOpts) []Clusters {
	var refinedClusters []Clusters

	for _, cluster := range clusters {
		if matchesFilters(cluster, opts) {
			refinedClusters = append(refinedClusters, cluster)
		}
	}

	return refinedClusters
}

func matchesFilters(cluster Clusters, opts ListOpts) bool {
	if opts.Name != "" && cluster.Metadata.Name != opts.Name {
		return false
	}
	if opts.ID != "" && cluster.Metadata.Id != opts.ID {
		return false
	}
	if opts.Type != "" && cluster.Spec.Type != opts.Type {
		return false
	}
	if opts.VpcID != "" && cluster.Spec.HostNetwork.VpcId != opts.VpcID {
		return false
	}
	if opts.Phase != "" && cluster.Status.Phase != opts.Phase {
		return false
	}
	return true
}
