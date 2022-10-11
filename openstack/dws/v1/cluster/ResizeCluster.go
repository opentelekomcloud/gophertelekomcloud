package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResizeClusterOpts struct {
	// Scale out an object.
	ScaleOut ScaleOut `json:"scale_out,omitempty"`
}

type ScaleOut struct {
	// Number of nodes to be added
	Count int32 `json:"count"`
}

func ResizeCluster(client *golangsdk.ServiceClient, clusterId string, opts ResizeClusterOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v1.0/{project_id}/clusters/{cluster_id}/resize
	_, err = client.Post(client.ServiceURL("clusters", clusterId, "resize"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
