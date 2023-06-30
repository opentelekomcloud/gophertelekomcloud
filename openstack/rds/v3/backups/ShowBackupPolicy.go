package backups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowBackupPolicy(client *golangsdk.ServiceClient, instanceId string) (*BackupPolicy, error) {
	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/backups/policy
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "backups", "policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res BackupPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "backup_policy")
	return &res, err
}

type BackupPolicy struct {
	// Indicates the number of days to retain the backup files.
	KeepDays int `json:"keep_days"`
	// Indicates the backup time window. Automated backups will be triggered during the backup time window.
	// The value must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format.
	// The HH value must be 1 greater than the hh value.
	// The values of mm and MM must be the same and must be set to any of the following: 00, 15, 30, or 45.
	StartTime string `json:"start_time,omitempty"`
	// Indicates the backup cycle configuration. Data will be automatically backed up on the selected days every week.
	Period string `json:"period,omitempty"`
}
