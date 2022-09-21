package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BackupInstanceBody struct {
	// Description of DCS instance backup.
	Remark string `json:"remark,omitempty"`
}

func BackupInstance(client *golangsdk.ServiceClient) (string, error) {
	raw, err := client.Post(client.ServiceURL("availableZones"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		BackupId string `json:"backup_id,omitempty"`
	}
	err = extract.Into(raw.Body, &res)
	return res.BackupId, err
}
