package clusters

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// RestartCluster function is used to restart a cluster.
func RestartCluster(client *golangsdk.ServiceClient, clusterID string) error {
	_, err := client.Post(client.ServiceURL("clusters", clusterID, "restart"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
