package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BackupInstanceOpts struct {
	// Description of DCS instance backup.
	Remark string `json:"remark,omitempty"`
}

func BackupInstance(client *golangsdk.ServiceClient, instancesId string, opts BackupInstanceOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("instances", instancesId, "backups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		BackupId string `json:"backup_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.BackupId, err
}
