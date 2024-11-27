package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts contains all the values needed to update a new cluster
type UpdateOpts struct {
	Spec UpdateSpec `json:"spec" required:"true"`
}

type UpdateSpec struct {
	// Cluster description
	Description string `json:"description,omitempty"`
}

// Update allows clusters to update description.
func Update(client *golangsdk.ServiceClient, clusterId string, opts UpdateOpts) (*Clusters, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /api/v3/projects/{project_id}/clusters/{cluster_id}
	raw, err := client.Put(client.ServiceURL("clusters", clusterId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Clusters
	return &res, extract.Into(raw.Body, &res)
}
