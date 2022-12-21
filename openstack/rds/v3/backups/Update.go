package backups

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// UpdateOpts contains all the values needed to update a Backup.
type UpdateOpts struct {
	InstanceId string `json:"-"`
	// Specifies the number of days to retain the generated backup files.
	// The value range is from 0 to 732. The value 0 indicates that the automated backup policy is disabled. To extend the retention period, contact customer service. Automated backups can be retained for up to 2562 days.
	// NOTICE
	// Once the automated backup policy is disabled, automated backups are no longer created and all incremental backups are deleted immediately. Operations related to the incremental backups, including downloads, replications, restorations, and rebuilds, may fail.
	// Disabling Automated Backup Policy is not allowed for SQL Server Primary/Standby and Cluster instances. So "keep_days" cannot be set to 0 for SQL Server Primary/Standby and Cluster instances.
	KeepDays int `json:"keep_days" required:"true"`
	// Specifies the backup time window. Automated backups will be triggered during the backup time window. This parameter is mandatory except that the automated backup policy is disabled.
	// The value must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format.
	// The HH value must be 1 greater than the hh value.
	// The values of mm and MM must be the same and must be set to any of the following: 00, 15, 30, or 45.
	// Example value:
	// 08:15-09:15
	// 23:00-00:00
	StartTime string `json:"start_time,omitempty"`
	// Specifies the backup cycle configuration. Data will be automatically backed up on the selected days every week. This parameter is mandatory except that the automated backup policy is disabled.
	// Value range: The value is digits separated by commas (,), indicating the day of the week and starting from Monday.
	// For example, the value 1,2,3,4 indicates that the backup period is Monday, Tuesday, Wednesday, and Thursday.
	Period string `json:"period,omitempty"`
}

// Update accepts a UpdateOpts struct and uses the values to update a Backup.The response code from api is 200
func Update(c *golangsdk.ServiceClient, opts UpdateOpts) (err error) {
	b, err := build.RequestBody(opts, "backup_policy")
	if err != nil {
		return
	}

	// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/backups/policy
	_, err = c.Put(c.ServiceURL("instances", opts.InstanceId, "backups", "policy"), b, nil,
		&golangsdk.RequestOpts{OkCodes: []int{200}, MoreHeaders: openstack.StdRequestOpts().MoreHeaders})
	return
}
