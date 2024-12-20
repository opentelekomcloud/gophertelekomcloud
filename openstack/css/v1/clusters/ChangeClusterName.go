package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ChangeClusterNameOpts struct {
	// DisplayName contains options for new name
	// This object is passed to the snapshots.ChangeClusterName function.
	DisplayName string `json:"displayName" required:"true"`
}

// ChangeClusterName function is used to change the name of a cluster.
func ChangeClusterName(client *golangsdk.ServiceClient, clusterID string, opts ChangeClusterNameOpts) error {
	// ChangeClusterName will change cluster name based on ChangeClusterNameOpts
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterID, "changename"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
