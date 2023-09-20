package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateBackupOpts struct {
	// Instance ID, which is compliant with the UUID format.
	InstanceId string `json:"instance_id"`
	// Backup name
	// The name consists of 4 to 64 characters and starts with a letter.
	// It is case-sensitive and can contain only letters, digits, hyphens (-), and underscores (_).
	Name string `json:"name"`
	// Backup description. It contains up to 256 characters and cannot contain the following special characters: >!<"&'=
	Description string `json:"description,omitempty"`
}

func CreateBackup(client *golangsdk.ServiceClient, opts CreateBackupOpts) (*CreateBackupResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/mysql/v3/{project_id}/backups/create
	raw, err := client.Post(client.ServiceURL("backups", "create"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res CreateBackupResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CreateBackupResponse struct {
	Backup Backup `json:"backup"`
	JobId  string `json:"job_id"`
}

type Backup struct {
	// Backup ID
	Id string `json:"id"`
	// Backup name
	Name string `json:"name"`
	// Backup description
	Description string `json:"description"`
	// Backup start time in the "yyyy-mm-ddThh:mm:ssZ" format,
	// where "T" indicates the start time of the time field, and "Z" indicates the time zone offset.
	BeginTime string `json:"begin_time"`
	// Backup status
	// Valid value:
	// BUILDING: Backup in progress
	// COMPLETED: Backup completed
	// FAILED: Backup failed
	// AVAILABLE: Backup available
	Status string `json:"status"`
	// Backup type
	// Valid value:
	// manual: manual full backup
	Type string `json:"type"`
	// Instance ID
	InstanceId string `json:"instance_id"`
}
