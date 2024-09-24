package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowBackupPolicy(client *golangsdk.ServiceClient, instanceId string) (*ShowBackupPolicyResult, error) {
	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/backups/policy
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "backups", "policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowBackupPolicyResult
	err = extract.IntoStructPtr(raw.Body, &res, "backup_policy")
	return &res, err
}

type ShowBackupPolicyResult struct {
	// Backup retention days
	KeepDays int32 `json:"keep_days"`
	// Backup time window. Automated backups will be triggered during the backup time window.
	StartTime string `json:"start_time"`
	// Backup cycle configuration. Data will be automatically backed up on the selected days every week.
	Period string `json:"period"`
}
