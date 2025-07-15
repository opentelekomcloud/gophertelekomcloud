package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RestoreToOriginalOpts struct {
	// Specifies the instance from which the backup was created
	Source Source `json:"source" required:"true"`
	// Specifies the instance to which the backup is restored.
	Target Target `json:"target" required:"true"`
}

type Source struct {
	// Specifies the instance ID, which can be obtained by calling the API for querying instances.
	InstanceId string `json:"instance_id,omitempty"`
	// Specifies the recovery mode. The enumerated values are as follows:
	// backup: indicates restoration from backup files. In this mode, backup_id is mandatory when type is optional.
	// timestamp: indicates point-in-time restoration. In this mode, restore_time is mandatory when type is mandatory.
	Type string `json:"type,omitempty"`
	// Specifies the ID of the backup to be restored.
	// This parameter must be specified when the backup file is used for restoration.
	BackupId string `json:"backup_id,omitempty"`
	// Specifies the time point of data restoration in the UNIX timestamp.
	// The unit is millisecond and the time zone is UTC.
	RestoreTime *int `json:"restore_time,omitempty"`
}

type Target struct {
	// Specifies ID of the DB instance to be restored from a backup.
	InstanceId string `json:"instance_id,omitempty"`
}

func RestoreToOriginal(client *golangsdk.ServiceClient, opts RestoreToOriginalOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/recovery
	raw, err := client.Post(client.ServiceURL("instances", "recovery"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	return ExtractJob(err, raw)
}
