package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type SetBackupPolicyOpts struct {
	InstanceId string
	// Backup policy object, which includes the backup retention period (days) and start time.
	BackupPolicy BackupPolicy `json:"backup_policy"`
}

type BackupPolicy struct {
	// Backup retention days. Value range: 0-35. The value 0 indicates that the automated backup policy is disabled.
	KeepDays int32 `json:"keep_days"`
	// Backup time window. Automated backups will be triggered during the backup time window.
	// This parameter is mandatory if the automated backup policy is enabled.
	// This parameter is not transferred if the automated backup policy is disabled.
	// The value must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format.
	// The HH value must be 1 greater than the hh value.
	// The values of mm and MM must be the same and must be set to any of the following: 00, 15, 30, and 45.
	// If this parameter is not transferred, the default backup time window is from 00:00 to 01:00.
	// Example value: 23:00-00:00
	StartTime string `json:"start_time,omitempty"`
	// Backup cycle configuration. Data will be automatically backed up on the selected days every week.
	// Value range: The value is a number separated by DBS case commas (,). The number indicates the week.
	// The restrictions on the backup retention period are as follows:
	// This parameter is not transferred if its value is set to 0.
	// If you set the retention period to 1 to 6 days, data is automatically backed up each day of the week.
	// Set the parameter value to 1,2,3,4,5,6,7.
	// If you set the retention period to 7 to 732 days, select at least one day of the week for the backup cycle.
	// Example value: 1,2,3,4
	Period string `json:"period,omitempty"`
}

func SetBackupPolicy(client *golangsdk.ServiceClient, opts SetBackupPolicyOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/backups/policy
	_, err = client.Put(client.ServiceURL("instances", opts.InstanceId, "backups", "policy"), b, nil, nil)
	return
}
