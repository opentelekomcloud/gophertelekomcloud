package policies

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/pointerto"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type BackupPolicy struct {
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
	ScheduledOperations []ScheduledOperation `json:"-"`
	// Backup policy status
	// disabled: indicates that the backup policy is unavailable.
	// enabled: indicates that the backup policy is available.
	Status string `json:"status"`
	// Tag list
	// Keys in the tag list must be unique.
	Tags []tags.ResourceTag `json:"tags"`
}

func extra(err error, raw *http.Response) (*BackupPolicy, error) {
	if err != nil {
		return nil, err
	}

	var res BackupPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "policy")
	return &res, err
}

type scheduledOperationString struct {
	Description         string                    `json:"description"`
	Enabled             bool                      `json:"enabled"`
	Name                string                    `json:"name"`
	OperationType       string                    `json:"operation_type"`
	OperationDefinition operationDefinitionString `json:"operation_definition"`
	Trigger             Trigger                   `json:"trigger" `
	ID                  string                    `json:"id"`
	TriggerID           string                    `json:"trigger_id"`
}

type operationDefinitionString struct {
	MaxBackups            interface{} `json:"max_backups"`
	RetentionDurationDays interface{} `json:"retention_duration_days"`
	Permanent             interface{} `json:"permanent"`
	PlanId                string      `json:"plan_id"`
	ProviderId            string      `json:"provider_id"`
	DayBackups            interface{} `json:"day_backups"`
	WeekBackups           interface{} `json:"week_backups"`
	MonthBackups          interface{} `json:"month_backups"`
	YearBackups           interface{} `json:"year_backups"`
	TimeZone              string      `json:"timezone"`
}

func toInt(v interface{}) int {
	var i int

	switch v := v.(type) {
	case string:
		i, _ = strconv.Atoi(v)
	case float64:
		i = int(v)
	}

	return i
}

// UnmarshalJSON helps to unmarshal BackupPolicy fields into needed values.
func (r *BackupPolicy) UnmarshalJSON(b []byte) error {
	type policy BackupPolicy
	var tmp struct {
		policy
		CreatedAt           golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
		ScheduledOperations []scheduledOperationString    `json:"scheduled_operations"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = BackupPolicy(tmp.policy)
	r.CreatedAt = time.Time(tmp.CreatedAt)

	for _, v := range tmp.ScheduledOperations {
		def := v.OperationDefinition

		var pt bool
		switch p := def.Permanent.(type) {
		case string:
			pt, _ = strconv.ParseBool(p)
		case bool:
			pt = p
		}

		r.ScheduledOperations = append(r.ScheduledOperations, ScheduledOperation{
			Description:   v.Description,
			Enabled:       v.Enabled,
			Name:          v.Name,
			OperationType: v.OperationType,
			OperationDefinition: OperationDefinition{
				MaxBackups:            pointerto.Int(toInt(def.MaxBackups)),
				RetentionDurationDays: toInt(def.RetentionDurationDays),
				Permanent:             pt,
				PlanId:                v.OperationDefinition.PlanId,
				ProviderId:            v.OperationDefinition.ProviderId,
				DayBackups:            toInt(def.DayBackups),
				WeekBackups:           toInt(def.WeekBackups),
				MonthBackups:          toInt(def.MonthBackups),
				YearBackups:           toInt(def.YearBackups),
				TimeZone:              v.OperationDefinition.TimeZone,
			},
			Trigger: Trigger{
				Properties: v.Trigger.Properties,
				Type:       v.Trigger.Type,
			},
			ID:        v.ID,
			TriggerID: v.TriggerID,
		})
	}

	return err
}
