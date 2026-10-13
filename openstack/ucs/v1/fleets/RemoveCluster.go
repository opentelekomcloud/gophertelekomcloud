package fleets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// RemoveCluster removes a cluster from a fleet via POST /v1/clusters/{clusterID}/unjoin.
func RemoveCluster(client *golangsdk.ServiceClient, clusterID string) error {
	_, err := client.Post(client.ServiceURL("clusters", clusterID, "unjoin"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
