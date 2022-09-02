package clusters

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
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
	raw, err := client.Get(client.ServiceURL("clusters"), nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: RequestOpts, JSONBody: nil,
	})
	if err != nil {
		return nil, err
	}

	var res ListCluster
	err = extract.Into(raw, &res)
	if err != nil {
		return nil, err
	}

	return filterClusters(res.Clusters, opts), nil
}

type ListCluster struct {
	// API type, fixed value Cluster
	Kind string `json:"kind"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion"`
	// all Clusters
	Clusters []Clusters `json:"items"`
}

func filterClusters(clusters []Clusters, opts ListOpts) []Clusters {
	var refinedClusters []Clusters
	var matched bool
	m := map[string]FilterStruct{}

	if opts.Name != "" {
		m["Name"] = FilterStruct{Value: opts.Name, Driller: []string{"Metadata"}}
	}
	if opts.ID != "" {
		m["Id"] = FilterStruct{Value: opts.ID, Driller: []string{"Metadata"}}
	}
	if opts.Type != "" {
		m["Type"] = FilterStruct{Value: opts.Type, Driller: []string{"Spec"}}
	}
	if opts.VpcID != "" {
		m["VpcId"] = FilterStruct{Value: opts.VpcID, Driller: []string{"Spec", "HostNetwork"}}
	}
	if opts.Phase != "" {
		m["Phase"] = FilterStruct{Value: opts.Phase, Driller: []string{"Status"}}
	}

	if len(m) > 0 && len(clusters) > 0 {
		for _, cluster := range clusters {
			matched = true

			for key, value := range m {
				if sVal := GetStructNestedField(&cluster, key, value.Driller); !(sVal == value.Value) {
					matched = false
				}
			}
			if matched {
				refinedClusters = append(refinedClusters, cluster)
			}
		}

	} else {
		refinedClusters = clusters
	}

	return refinedClusters
}
