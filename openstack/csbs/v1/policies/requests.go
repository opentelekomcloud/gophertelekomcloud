package policies

type ListOpts struct {
	ID         string `json:"id"`
	Name       string `q:"name"`
	Sort       string `q:"sort"`
	Status     string `json:"status"`
	Limit      int    `q:"limit"`
	Marker     string `q:"marker"`
	Offset     int    `q:"offset"`
	AllTenants string `q:"all_tenants"`
}

type PolicyParam struct {
	Common interface{} `json:"common,omitempty"`
}

type Resource struct {
	Id        string      `json:"id" required:"true"`
	Type      string      `json:"type" required:"true"`
	Name      string      `json:"name" required:"true"`
	ExtraInfo interface{} `json:"extra_info,omitempty"`
}

type ScheduledOperation struct {
	Description         string              `json:"description,omitempty"`
	Enabled             bool                `json:"enabled"`
	Name                string              `json:"name,omitempty"`
	OperationType       string              `json:"operation_type" required:"true"`
	OperationDefinition OperationDefinition `json:"operation_definition" required:"true"`
	Trigger             Trigger             `json:"trigger" required:"true"`
}

type OperationDefinition struct {
	MaxBackups            int    `json:"max_backups,omitempty"`
	RetentionDurationDays int    `json:"retention_duration_days,omitempty"`
	Permanent             bool   `json:"permanent"`
	PlanId                string `json:"plan_id,omitempty"`
	ProviderId            string `json:"provider_id,omitempty"`
	DayBackups            int    `json:"day_backups,omitempty"`
	WeekBackups           int    `json:"week_backups,omitempty"`
	MonthBackups          int    `json:"month_backups,omitempty"`
	YearBackups           int    `json:"year_backups,omitempty"`
	TimeZone              string `json:"timezone,omitempty"`
}

type Trigger struct {
	Properties TriggerProperties `json:"properties" required:"true"`
}

type TriggerProperties struct {
	Pattern string `json:"pattern" required:"true"`
}

type ResourceTag struct {
	Key   string `json:"key" required:"true"`
	Value string `json:"value" required:"true"`
}

type ScheduledOperationToUpdate struct {
	Description         string              `json:"description,omitempty"`
	Enabled             bool                `json:"enabled"`
	TriggerId           string              `json:"trigger_id,omitempty"`
	Name                string              `json:"name,omitempty"`
	OperationDefinition OperationDefinition `json:"operation_definition,omitempty"`
	Trigger             Trigger             `json:"trigger,omitempty"`
	Id                  string              `json:"id" required:"true"`
}
