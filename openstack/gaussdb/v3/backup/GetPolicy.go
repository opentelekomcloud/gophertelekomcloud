package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetPolicy(client *golangsdk.ServiceClient, instanceId string) (*BackupPolicy, error) {
	// GET https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}/backups/policy
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "backups", "policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res BackupPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "backup_policy")
	return &res, err
}

type BackupPolicy struct {
	// Backup retention days
	KeepDays int32 `json:"keep_days"`
	// Backup time window. The creation of an automated backup will be triggered during the backup time window.
	StartTime string `json:"start_time"`
	// Backup cycle configuration. Data will be automatically backed up on the selected days every week.
	Period string `json:"period"`
	// Number of retained level-1 backups. This parameter is returned when level-1 backup is enabled
	RetentionNumBackupLevel1 int32 `json:"retention_num_backup_level1"`
}
