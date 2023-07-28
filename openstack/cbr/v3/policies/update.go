package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Enabled             *bool           `json:"enabled,omitempty"`
	Name                string          `json:"name,omitempty"`
	OperationDefinition *PolicyODCreate `json:"operation_definition,omitempty"`
	Trigger             *Trigger        `json:"trigger,omitempty"`
}

// PolicyODCreate is policy operation definition
// see https://docs.otc.t-systems.com/en-us/api/cbr/CreatePolicy.html#CreatePolicy__request_PolicyoODCreate
type PolicyODCreate struct {
	DailyBackups          int    `json:"day_backups"`
	WeekBackups           int    `json:"week_backups"`
	YearBackups           int    `json:"year_backups"`
	MonthBackups          int    `json:"month_backups"`
	MaxBackups            int    `json:"max_backups,omitempty"`
	RetentionDurationDays int    `json:"retention_duration_days,omitempty"`
	Timezone              string `json:"timezone,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Policy, error) {
	b, err := build.RequestBodyMap(opts, "policy")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("policies", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Policy
	err = extract.IntoStructPtr(raw.Body, &res, "policy")
	return &res, err
}
