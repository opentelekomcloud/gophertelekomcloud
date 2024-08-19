package backups

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, backupId string) (*Job, error) {
	// DELETE https://{Endpoint}/v3/{project_id}/backups/{backup_id}
	raw, err := client.Delete(client.ServiceURL("backups", backupId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extractJob(err, raw)
}
