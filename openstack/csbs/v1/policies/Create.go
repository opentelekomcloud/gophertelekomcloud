package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateOpts contains the options for create a Backup Policy. This object is passed to policies.Create().
type CreateOpts struct {
	// Backup policy description
	// The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	Description string `json:"description,omitempty"`
	// Backup policy name
	// The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name" required:"true"`
	// Backup parameters
	// For details, see Table 2-24.
	Parameters PolicyParam `json:"parameters" required:"true"`
	// Backup provider ID, which specifies whether the backup object is a server or disk.
	// This parameter has a fixed value. For CSBS, the value is fc4d5750-22e7-4798-8a46-f48f62c4c1da.
	ProviderId string `json:"provider_id" required:"true"`
	// Backup object list. The list can be blank.
	// For details, see Table 2-25.
	Resources []Resource `json:"resources" required:"true"`
	// Scheduling period
	ScheduledOperations []ScheduledOperation `json:"scheduled_operations" required:"true"`
	// Tag list
	// This list cannot be an empty list.
	// The list can contain up to 10 keys.
	// Keys in this list must be unique.
	Tags []tags.ResourceTag `json:"tags,omitempty"`
}

type PolicyParam struct {
	// General backup policy parameters, which are blank by default
	Common interface{} `json:"common,omitempty"`
}

type ScheduledOperation struct {
	// Scheduling period description
	// The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	Description string `json:"description,omitempty"`
	// Whether the backup policy is enabled
	// If it is set to true, automatic scheduling is enabled. If it is set to false, automatic scheduling is disabled but you can execute the policy manually.
	Enabled bool `json:"enabled"`
	// Whether the backup policy is enabled
	// If it is set to true, automatic scheduling is enabled. If it is set to false, automatic scheduling is disabled but you can execute the policy manually.
	Name string `json:"name,omitempty"`
	// Operation type, which can be backup
	// Enumeration values: backup
	OperationType string `json:"operation_type" required:"true"`
	// Scheduling period parameters
	OperationDefinition OperationDefinition `json:"operation_definition" required:"true"`
	// Scheduling policy
	Trigger Trigger `json:"trigger" required:"true"`
	// Response: Scheduling period ID
	ID string `json:"id,omitempty"`
	// Response: Scheduler ID
	TriggerID string `json:"trigger_id,omitempty"`
}

type OperationDefinition struct {
	// Maximum number of backups that can be automatically created for a backup object.
	// The value can be -1 or ranges from 0 to 99999. If the value is set to -1,
	// the backups will not be cleared even though the configured retained backup quantity limit is exceeded.
	MaxBackups int `json:"max_backups,omitempty"`
	// Duration of retaining a backup, in days. The value can be -1 or ranges from 0 to 99999.
	// If the value is set to -1, backups will not be cleared even though the configured retention duration is exceeded.
	RetentionDurationDays int `json:"retention_duration_days,omitempty"`
	// Whether backups are permanently retained. false: no. true: yes
	Permanent bool `json:"permanent"`
	// Backup policy ID
	PlanId string `json:"plan_id,omitempty"`
	// Backup provider ID, which specifies whether the backup object is a server or disk.
	// This parameter has a fixed value. For CSBS, the value is fc4d5750-22e7-4798-8a46-f48f62c4c1da.
	ProviderId string `json:"provider_id,omitempty"`
	// Specifies the maximum number of retained daily backups. The latest backup of each day is saved in the long term.
	// This parameter can be effective together with the maximum number of retained backups specified by max_backups.
	// If this parameter is configured, timezone is mandatory.
	DayBackups int `json:"day_backups,omitempty"`
	// Specifies the maximum number of retained weekly backups. The latest backup of each week is saved in the long term.
	// This parameter can be effective together with the maximum number of retained backups specified by max_backups.
	// If this parameter is configured, timezone is mandatory.
	WeekBackups int `json:"week_backups,omitempty"`
	// Specifies the maximum number of retained monthly backups. The latest backup of each month is saved in the long term.
	// This parameter can be effective together with the maximum number of retained backups specified by max_backups.
	// If this parameter is configured, timezone is mandatory.
	MonthBackups int `json:"month_backups,omitempty"`
	// Specifies the maximum number of retained yearly backups. The latest backup of each year is saved in the long term.
	// This parameter can be effective together with the maximum number of retained backups specified by max_backups.
	// If this parameter is configured, timezone is mandatory.
	YearBackups int `json:"year_backups,omitempty"`
	// Time zone where the user is located, for example, UTC+08:00. Set this parameter only
	// after you have configured any of the parameters day_backups, week_backups, month_backups, and year_backups.
	TimeZone string `json:"timezone,omitempty"`
}

type Trigger struct {
	// Scheduler properties
	Properties TriggerProperties `json:"properties" required:"true"`
	// Response: Scheduler name
	Name string `json:"name,omitempty"`
	// Response: Scheduler ID
	ID string `json:"id,omitempty"`
	// Response: Scheduling type. The value is fixed at time.
	Type string `json:"type,omitempty"`
}

type TriggerProperties struct {
	// Scheduling policy of the scheduler. The value consists of a maximum of 10,240 characters.
	// The scheduling policy complies with iCalendar RFC 2445, but it supports only four parameters,
	// which are FREQ, BYDAY, BYHOUR, and BYMINUTE. FREQ can be set only to WEEKLY or DAILY.
	// BYDAY can be set to MO, TU, WE, TH, FR, SA, or SU (seven days of a week).
	// BYHOUR ranges from 0 to 23 hours. BYMINUTE ranges from 0 to 59 minutes.
	// The scheduling interval cannot be less than 1 hour. A maximum of 24 time points are allowed in a day.
	Pattern string `json:"pattern" required:"true"`
	// Response: Scheduler start time, for example, 2017-04-18T01:21:52
	StartTime string `json:"start_time,omitempty"`
	// Response: Scheduler type
	Format string `json:"format,omitempty"`
}

// Create will create a new backup policy based on the values in CreateOpts. To extract
// the Backup object from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*BackupPolicy, error) {
	b, err := build.RequestBody(opts, "policy")
	if err != nil {
		return nil, err
	}

	// POST https://{endpoint}/v1/{project_id}/policies
	raw, err := client.Post(client.ServiceURL("policies"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
