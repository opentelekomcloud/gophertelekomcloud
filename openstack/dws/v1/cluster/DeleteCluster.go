package cluster

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

type DeleteClusterOpts struct {
	ClusterId string
	// The number of the latest manual snapshots that need to be retained for a cluster.
	KeepLastManualSnapshot *int `json:"keep_last_manual_snapshot" required:"true"`
}

func DeleteCluster(client *golangsdk.ServiceClient, opts DeleteClusterOpts) error {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return err
	}

	// DELETE /v1.0/{project_id}/clusters/{cluster_id}
	_, err = client.Delete(client.ServiceURL("clusters", opts.ClusterId)+q.String(), nil)
	return err
}
