package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func DeleteBackupFile(client *golangsdk.ServiceClient, instancesId string, backupId string) (err error) {
	_, err = client.Delete(client.ServiceURL("instances", instancesId, "backups", backupId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
