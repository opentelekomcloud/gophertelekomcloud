package clusters

import (
	"strings"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// ScaleInOpts defines options for scaling in a cluster.
type ScaleInOpts struct {
	// Type specifies the type of instance to be scaled in.
	// Select at least one from `ess`, `ess-cold`, `ess-master`, and `ess-client`.
	Type string `json:"type"`
	// ReduceNodeNum is the number of nodes to be removed.
	// After scaling in, there must be at least one node in each AZ under each node type.
	// In a cross-AZ cluster, the difference between the number of nodes of the same type in different AZs cannot exceed 1.
	// For a cluster with Master nodes, the number of removed master nodes in a scale-in must be fewer than half of the original master node count.
	ReduceNodeNum int `json:"reducedNodeNum"`
}

// ScaleInRequest is a wrapper to structure the "shrink" key in the JSON body.
type ScaleInRequest struct {
	Shrink []ScaleInOpts `json:"shrink"`
}

// ScaleInCluster scales in a cluster by removing specified nodes.
func ScaleInCluster(client *golangsdk.ServiceClient, clusterID string, opts []ScaleInOpts) error {
	// Wrap opts in ScaleInRequest to match the required JSON structure.
	request := ScaleInRequest{
		Shrink: opts,
	}

	b, err := build.RequestBody(request, "")
	if err != nil {
		return err
	}

	// POST /v1.0/extend/{project_id}/clusters/{cluster_id}/role/shrink
	url := client.ServiceURL("clusters", clusterID, "role", "shrink")
	convertedURL := strings.Replace(url, "v1.0", "v1.0/extend", 1)

	_, err = client.Post(convertedURL, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
