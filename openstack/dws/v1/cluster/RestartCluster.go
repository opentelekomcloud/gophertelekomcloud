package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RestartClusterOpts struct {
	// Restart flag.
	Restart interface{} `json:"restart"`
}

func RestartCluster(client *golangsdk.ServiceClient, clusterId string, opts RestartClusterOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v1.0/{project_id}/clusters/{cluster_id}/restart
	_, err = client.Post(client.ServiceURL("clusters", clusterId, "restart"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
