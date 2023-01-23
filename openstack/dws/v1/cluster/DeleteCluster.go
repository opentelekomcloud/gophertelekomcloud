package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteClusterOpts struct {
	ClusterId string `json:"-"`
	// The number of the latest manual snapshots that need to be retained for a cluster.
	KeepLastManualSnapshot *int `json:"keep_last_manual_snapshot" required:"true"`
}

func DeleteCluster(client *golangsdk.ServiceClient, opts DeleteClusterOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// DELETE /v1.0/{project_id}/clusters/{cluster_id}
	_, err = client.DeleteWithBody(client.ServiceURL("clusters", opts.ClusterId), b, nil)
	return err
}
