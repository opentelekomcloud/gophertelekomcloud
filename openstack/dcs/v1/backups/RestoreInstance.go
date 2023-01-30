package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type RestoreInstanceOpts struct {
	// ID of the backup record
	BackupId string `json:"backup_id"`
	// Description of DCS instance restoration
	Remark string `json:"remark,omitempty"`
}

func RestoreInstance(client *golangsdk.ServiceClient, instancesId string, opts RestoreInstanceOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("instances", instancesId, "restores"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		RestoreId string `json:"restore_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.RestoreId, err
}
