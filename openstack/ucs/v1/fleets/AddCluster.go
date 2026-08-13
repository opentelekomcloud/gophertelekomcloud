package fleets

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type AddClusterOpts struct {
	ClusterGroupID string `json:"clusterGroupID" required:"true"`
}

// AddCluster adds a cluster to a fleet via POST /v1/clusters/{clusterID}/join.
func AddCluster(client *golangsdk.ServiceClient, clusterID string, opts AddClusterOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterID, "join"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
