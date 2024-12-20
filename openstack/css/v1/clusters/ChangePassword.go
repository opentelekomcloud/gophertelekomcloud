package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ChangePasswordOpts struct {
	// DisplayName contains options for new name
	// This object is passed to the snapshots.ChangeClusterName function.
	NewPassword string `json:"newpassword" required:"true"`
}

// ChangePassword function is used to change the password of a cluster.
func ChangePassword(client *golangsdk.ServiceClient, clusterID string, opts ChangePasswordOpts) (err error) {
	// ChangeClusterName will change cluster name based on ChangeClusterNameOpts
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterID, "password", "reset"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
