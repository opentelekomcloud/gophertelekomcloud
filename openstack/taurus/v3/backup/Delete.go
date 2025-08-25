package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Delete(client *golangsdk.ServiceClient, backupID string) (*DeleteResponse, error) {
	raw, err := client.Delete(client.ServiceURL("backups", backupID), &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	return &res, extract.Into(raw.Body, &res)
}

type DeleteResponse struct {
	BackupId   string `json:"backup_id"`
	BackupName string `json:"backup_name"`
}
