package policies

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateBackupPolicy struct {
	CreatedAt           time.Time                      `json:"-"`
	Description         string                         `json:"description"`
	ID                  string                         `json:"id"`
	Name                string                         `json:"name"`
	Parameters          PolicyParam                    `json:"parameters"`
	ProjectId           string                         `json:"project_id"`
	ProviderId          string                         `json:"provider_id"`
	Resources           []Resource                     `json:"resources"`
	ScheduledOperations []CreateScheduledOperationResp `json:"scheduled_operations"`
	Status              string                         `json:"status"`
	Tags                []ResourceTag                  `json:"tags"`
}

func extra(err error, raw *http.Response) (*CreateBackupPolicy, error) {
	if err != nil {
		return nil, err
	}

	var res CreateBackupPolicy
	err = extract.IntoSlicePtr(raw.Body, &res, "policy")
	return &res, err
}

type BackupPolicy struct {
	CreatedAt           time.Time                `json:"-"`
	Description         string                   `json:"description"`
	ID                  string                   `json:"id"`
	Name                string                   `json:"name"`
	Parameters          PolicyParam              `json:"parameters"`
	ProjectId           string                   `json:"project_id"`
	ProviderId          string                   `json:"provider_id"`
	Resources           []Resource               `json:"resources"`
	ScheduledOperations []ScheduledOperationResp `json:"scheduled_operations"`
	Status              string                   `json:"status"`
	Tags                []ResourceTag            `json:"tags"`
}

type ScheduledOperationResp struct {
	Description         string                  `json:"description"`
	Enabled             bool                    `json:"enabled"`
	Name                string                  `json:"name"`
	OperationType       string                  `json:"operation_type"`
	OperationDefinition OperationDefinitionResp `json:"operation_definition"`
	Trigger             TriggerResp             `json:"trigger"`
	ID                  string                  `json:"id"`
	TriggerID           string                  `json:"trigger_id"`
}

type CreateScheduledOperationResp struct {
	Description         string                        `json:"description"`
	Enabled             bool                          `json:"enabled"`
	Name                string                        `json:"name"`
	OperationType       string                        `json:"operation_type"`
	OperationDefinition CreateOperationDefinitionResp `json:"operation_definition"`
	Trigger             TriggerResp                   `json:"trigger"`
	ID                  string                        `json:"id"`
	TriggerID           string                        `json:"trigger_id"`
}

type OperationDefinitionResp struct {
	MaxBackups            int    `json:"max_backups"`
	RetentionDurationDays int    `json:"retention_duration_days"`
	Permanent             bool   `json:"permanent"`
	PlanId                string `json:"plan_id"`
	ProviderId            string `json:"provider_id"`
}

type CreateOperationDefinitionResp struct {
	MaxBackups            int    `json:"-"`
	RetentionDurationDays int    `json:"-"`
	Permanent             bool   `json:"-"`
	PlanId                string `json:"plan_id"`
	ProviderId            string `json:"provider_id"`
}

type TriggerResp struct {
	Properties TriggerPropertiesResp `json:"properties"`
	Name       string                `json:"name"`
	ID         string                `json:"id"`
	Type       string                `json:"type"`
}

type TriggerPropertiesResp struct {
	Pattern   string    `json:"pattern"`
	StartTime time.Time `json:"-"`
}

// UnmarshalJSON helps to unmarshal BackupPolicy fields into needed values.
func (r *BackupPolicy) UnmarshalJSON(b []byte) error {
	type tmp BackupPolicy
	var s struct {
		tmp
		CreatedAt golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = BackupPolicy(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)

	return err
}

// UnmarshalJSON helps to unmarshal TriggerPropertiesResp fields into needed values.
func (r *TriggerPropertiesResp) UnmarshalJSON(b []byte) error {
	type tmp TriggerPropertiesResp
	var s struct {
		tmp
		StartTime golangsdk.JSONRFC3339ZNoTNoZ `json:"start_time"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = TriggerPropertiesResp(s.tmp)

	r.StartTime = time.Time(s.StartTime)

	return err
}

// UnmarshalJSON helps to unmarshal OperationDefinitionResp fields into needed values.
func (r *CreateOperationDefinitionResp) UnmarshalJSON(b []byte) error {
	type tmp CreateOperationDefinitionResp
	var s struct {
		tmp
		MaxBackups            string `json:"max_backups"`
		RetentionDurationDays string `json:"retention_duration_days"`
		Permanent             string `json:"permanent"`
	}

	err := json.Unmarshal(b, &s)

	if err != nil {
		switch err.(type) {
		case *json.UnmarshalTypeError: // check if type error occurred (handles if no type conversion is required for cloud)

			var s struct {
				tmp
				MaxBackups            int  `json:"max_backups"`
				RetentionDurationDays int  `json:"retention_duration_days"`
				Permanent             bool `json:"permanent"`
			}
			err := json.Unmarshal(b, &s)
			if err != nil {
				return err
			}
			*r = CreateOperationDefinitionResp(s.tmp)
			r.MaxBackups = s.MaxBackups
			r.RetentionDurationDays = s.RetentionDurationDays
			r.Permanent = s.Permanent
			return nil
		default:
			return err
		}
	}

	*r = CreateOperationDefinitionResp(s.tmp)

	switch s.MaxBackups {
	case "":
		r.MaxBackups = 0
	default:
		r.MaxBackups, err = strconv.Atoi(s.MaxBackups)
		if err != nil {
			return err
		}
	}

	switch s.RetentionDurationDays {
	case "":
		r.RetentionDurationDays = 0
	default:
		r.RetentionDurationDays, err = strconv.Atoi(s.RetentionDurationDays)
		if err != nil {
			return err
		}
	}

	switch s.Permanent {
	case "":
		r.Permanent = false
	default:
		r.Permanent, err = strconv.ParseBool(s.Permanent)
		if err != nil {
			return err
		}
	}

	return err
}
