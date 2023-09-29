package backup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePolicyOpts struct {
	// Instance ID, which is compliant with the UUID format.
	InstanceId   string
	BackupPolicy UpdateBackupPolicy `json:"backup_policy"`
}

type UpdateBackupPolicy struct {
	// Backup time window. The creation of an automated backup will be triggered during the backup time window. The value cannot be empty.
	// It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format.
	// The HH value must be 1 greater than the hh value.
	// The values of mm and MM must be the same and must be set to 00. Example value: 21:00-22:00
	StartTime string `json:"start_time"`
	// Backup retention days
	KeepDays int `json:"keep_days"`
	// Backup cycle configuration. Data will be automatically backed up on the selected days every week. Value range:
	// The value is a number separated by commas (,), indicating the days of the week.
	// For example, the value 1,2,3,4 indicates that the backup period is every Monday, Tuesday,Wednesday, and Thursday.
	Period string `json:"period"`
}

func UpdatePolicy(client *golangsdk.ServiceClient, opts UpdatePolicyOpts) (*UpdatepPolicyResponse, error) {
	// PUT https://{Endpoint}/mysql/v3/{project_id}/instances/{instance_id}/backups/policy/update
	b, err := build.RequestBody(opts.BackupPolicy, "backup_policy")
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(client.ServiceURL("instances", opts.InstanceId, "backups", "policy", "update"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 201},
	})
	if err != nil {
		return nil, err
	}

	var res UpdatepPolicyResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdatepPolicyResponse struct {
	// Backup status. Value:
	// BUILDING: Modification in progress
	// COMPLETED: Modification completed
	// FAILED: Modification failed
	Status string `json:"status"`
	// Instance ID
	InstanceId string `json:"instance_id"`
	// Instance name
	InstanceName string `json:"instance_name"`
}
