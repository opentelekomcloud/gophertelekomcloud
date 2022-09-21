package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DeleteBackupFile(client *golangsdk.ServiceClient) (string, error) {
	raw, err := client.Post(client.ServiceURL("availableZones"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return res.BackupId, err
}
