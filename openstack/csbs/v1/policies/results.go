package policies

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type BackupPolicyString struct {
	// Creation time, for example, 2017-04-18T01:21:52.701973
	CreatedAt time.Time `json:"-"`
	// Backup policy description
	// The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	Description string `json:"description"`
	// Backup policy ID
	ID string `json:"id"`
	// Backup policy name
	// The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).
	Name string `json:"name"`
	// Parameters of a backup policy
	Parameters PolicyParam `json:"parameters"`
	// Project ID
	ProjectId string `json:"project_id"`
	// Backup provider ID, which specifies whether the backup object is a server or disk.
	// This parameter has a fixed value. For CSBS, the value is fc4d5750-22e7-4798-8a46-f48f62c4c1da.
	ProviderId string `json:"provider_id"`
	// Backup object list
	Resources []Resource `json:"resources"`
	// Scheduling period list
	ScheduledOperations []ScheduledOperationString `json:"scheduled_operations"`
	// Backup policy status
	// disabled: indicates that the backup policy is unavailable.
	// enabled: indicates that the backup policy is available.
	Status string `json:"status"`
	// Tag list
	// Keys in the tag list must be unique.
	Tags []tags.ResourceTag `json:"tags"`
}

func extra(err error, raw *http.Response) (*BackupPolicyString, error) {
	if err != nil {
		return nil, err
	}

	var res BackupPolicyString
	err = extract.IntoStructPtr(raw.Body, &res, "policy")
	return &res, err
}

type ScheduledOperationString struct {
	// Scheduling period description
	// The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	Description string `json:"description"`
	// Whether the backup policy is enabled
	// If it is set to true, automatic scheduling is enabled. If it is set to false, automatic scheduling is disabled but you can execute the policy manually.
	Enabled bool `json:"enabled"`
	// Whether the backup policy is enabled
	// If it is set to true, automatic scheduling is enabled. If it is set to false, automatic scheduling is disabled but you can execute the policy manually.
	Name string `json:"name"`
	// Operation type, which can be backup
	// Enumeration values: backup
	OperationType string `json:"operation_type"`
	// Scheduling period parameters
	OperationDefinition OperationDefinitionString `json:"operation_definition"`
	// Scheduling policy
	Trigger Trigger `json:"trigger" `
	// Response: Scheduling period ID
	ID string `json:"id"`
	// Response: Scheduler ID
	TriggerID string `json:"trigger_id"`
}

type OperationDefinitionString struct {
	// Maximum number of backups that can be automatically created for a backup object.
	// The value can be -1 or ranges from 0 to 99999. If the value is set to -1,
	// the backups will not be cleared even though the configured retained backup quantity limit is exceeded.
	MaxBackups string `json:"max_backups"`
	// Duration of retaining a backup, in days. The value can be -1 or ranges from 0 to 99999.
	// If the value is set to -1, backups will not be cleared even though the configured retention duration is exceeded.
	RetentionDurationDays string `json:"retention_duration_days"`
	// Whether backups are permanently retained. false: no. true: yes
	Permanent string `json:"permanent"`
	// Backup policy ID
	PlanId string `json:"plan_id"`
	// Backup provider ID, which specifies whether the backup object is a server or disk.
	// This parameter has a fixed value. For CSBS, the value is fc4d5750-22e7-4798-8a46-f48f62c4c1da.
	ProviderId string `json:"provider_id"`
	// Specifies the maximum number of retained daily backups. The latest backup of each day is saved in the long term.
	// This parameter can be effective together with the maximum number of retained backups specified by max_backups.
	// If this parameter is configured, timezone is mandatory.
	DayBackups string `json:"day_backups"`
	// Specifies the maximum number of retained weekly backups. The latest backup of each week is saved in the long term.
	// This parameter can be effective together with the maximum number of retained backups specified by max_backups.
	// If this parameter is configured, timezone is mandatory.
	WeekBackups string `json:"week_backups"`
	// Specifies the maximum number of retained monthly backups. The latest backup of each month is saved in the long term.
	// This parameter can be effective together with the maximum number of retained backups specified by max_backups.
	// If this parameter is configured, timezone is mandatory.
	MonthBackups string `json:"month_backups"`
	// Specifies the maximum number of retained yearly backups. The latest backup of each year is saved in the long term.
	// This parameter can be effective together with the maximum number of retained backups specified by max_backups.
	// If this parameter is configured, timezone is mandatory.
	YearBackups string `json:"year_backups"`
	// Time zone where the user is located, for example, UTC+08:00. Set this parameter only
	// after you have configured any of the parameters day_backups, week_backups, month_backups, and year_backups.
	TimeZone string `json:"timezone"`
}

// UnmarshalJSON helps to unmarshal BackupPolicyString fields into needed values.
func (r *BackupPolicyString) UnmarshalJSON(b []byte) error {
	type tmp BackupPolicyString
	var s struct {
		tmp
		CreatedAt golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = BackupPolicyString(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)

	return err
}
