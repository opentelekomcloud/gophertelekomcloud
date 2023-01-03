package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResizeClusterOpts struct {
	ClusterId string `json:"-"`
	// Number of nodes to be added
	Count int `json:"count" required:"true"`
}

func ResizeCluster(client *golangsdk.ServiceClient, opts ResizeClusterOpts) (err error) {
	b, err := build.RequestBody(opts, "scale_out")
	if err != nil {
		return
	}

	// POST /v1.0/{project_id}/clusters/{cluster_id}/resize
	_, err = client.Post(client.ServiceURL("clusters", opts.ClusterId, "resize"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
