package snapshots

import "github.com/opentelekomcloud/gophertelekomcloud"

type UpdateConfigurationOpts struct {
	// OBS bucket used for index data backup.
	// If there is snapshot data in an OBS bucket, only the OBS bucket is used and cannot be changed.
	Bucket string `json:"bucket" required:"true"`
	// IAM agency used to access OBS.
	Agency string `json:"agency" required:"true"`
	// Key ID used for snapshot encryption.
	SnapshotCmkID string `json:"snapshotCmkId,omitempty"`
}

func UpdateConfiguration(client *golangsdk.ServiceClient, clusterID string, opts UpdateConfigurationOpts) (err error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterID, "index_snapshot", "setting"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
