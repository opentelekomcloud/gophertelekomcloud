package logs

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type EnableAutomaticBackupOpts struct {
	// This parameter passed to the logs.EnableAutomaticBackups function.
	// Period is the start time of a backup job.
	Period string `json:"period" required:"true"`
}

// EnableAutomaticBackups will enable the automatic log backup policy for the cluster based on EnableAutomaticBackupOpts.
func EnableAutomaticBackups(client *golangsdk.ServiceClient, clusterID string, opts EnableAutomaticBackupOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("clusters", clusterID, "logs", "policy", "update")

	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return err
}
