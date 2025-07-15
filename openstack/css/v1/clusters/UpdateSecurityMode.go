package clusters

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type SecurityModeOpts struct {
	// Indicates whether to enable the security mode.
	AuthorityEnabled *bool `json:"authorityEnable"`
	// Cluster password.
	AdminPassword string `json:"adminPwd"`
	// Indicates whether to enable HTTPS.
	HttpsEnabled *bool `json:"httpsEnable"`
}

// UpdateSecurityMode - change the security mode of a cluster.
func UpdateSecurityMode(client *golangsdk.ServiceClient, clusterID string, opts SecurityModeOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "mode", "change")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
