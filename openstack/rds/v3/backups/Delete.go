package backups

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(c *golangsdk.ServiceClient, backupId string) (err error) {
	// DELETE https://{Endpoint}/v3/{project_id}/backups/{backup_id}
	_, err = c.Delete(c.ServiceURL("backups", backupId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 201, 204},
	})
	return
}
