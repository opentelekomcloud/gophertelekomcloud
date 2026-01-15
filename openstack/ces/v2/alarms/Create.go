package alarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains the options for creating an alarm rule.
type CreateOpts struct {
	// Specifies the alarm rule name.
	// Enter 1 to 128 characters. Only letters, digits, underscores (_), and hyphens (-) are allowed.
	Name string `json:"name" required:"true"`
	// Provides supplementary information about the alarm rule. Enter 0 to 256 characters.
	Description string `json:"description,omitempty"`
	// Specifies the namespace of a service.
	// The value must be in the service.item format and can contain 3 to 32 characters.
	Namespace string `json:"namespace" required:"true"`
	// Specifies the resource group ID.
	// This parameter is mandatory when type is set to RESOURCE_GROUP.
	ResourceGroupId string `json:"resource_group_id,omitempty"`
	// Specifies the resource list. This parameter is mandatory when type is set to MULTI_INSTANCE.
	// A maximum of 1000 resources are supported.
	Resources [][]Dimension `json:"resources,omitempty"`
	// Specifies the alarm policies.
	// Either policies or alarm_template_id must be configured.
	Policies []Policy `json:"policies,omitempty"`
	// Specifies the alarm rule type.
	// ALL_INSTANCE: all resources of a service
	// RESOURCE_GROUP: resource group
	// MULTI_INSTANCE: specified resources
	// EVENT.SYS: system event
	// EVENT.CUSTOM: custom event
	Type string `json:"type" required:"true"`
	// Specifies the alarm template ID. Either policies or alarm_template_id must be configured.
	AlarmTemplateId string `json:"alarm_template_id,omitempty"`
	// Specifies whether to generate alarms.
	Enabled *bool `json:"enabled,omitempty"`
	// Specifies whether to enable notification.
	NotificationEnabled *bool `json:"notification_enabled,omitempty"`
	// Specifies the action to be triggered when an alarm is generated.
	AlarmNotifications []Notification `json:"alarm_notifications,omitempty"`
	// Specifies the action to be triggered after an alarm is cleared.
	OkNotifications []Notification `json:"ok_notifications,omitempty"`
	// Specifies the time when the alarm notification was enabled. The value is in the format of HH:MM.
	NotificationBeginTime string `json:"notification_begin_time,omitempty"`
	// Specifies the time when the alarm notification was disabled. The value is in the format of HH:MM.
	NotificationEndTime string `json:"notification_end_time,omitempty"`
	// Specifies the enterprise project ID.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

// Dimension represents a metric dimension.
type Dimension struct {
	// Specifies the dimension name.
	Name string `json:"name" required:"true"`
	// Specifies the dimension value.
	Value string `json:"value,omitempty"`
}

// Policy represents an alarm policy.
type Policy struct {
	// Specifies the metric name.
	MetricName string `json:"metric_name" required:"true"`
	// Specifies the rollup period in seconds.
	// Possible values: 0, 1, 300, 1200, 3600, 14400, 86400
	Period int `json:"period" required:"true"`
	// Specifies the data rollup method.
	// Possible values: average, max, min, sum, variance
	Filter string `json:"filter" required:"true"`
	// Specifies the comparison operator.
	// Possible values: >, <, >=, <=, =, !=
	ComparisonOperator string `json:"comparison_operator" required:"true"`
	// Specifies the alarm threshold.
	Value float64 `json:"value" required:"true"`
	// Specifies the data unit.
	Unit string `json:"unit,omitempty"`
	// Specifies the number of consecutive times that the alarm condition is met.
	Count int `json:"count" required:"true"`
	// Specifies the interval for triggering an alarm if the alarm persists in seconds.
	// Possible values: 0, 300, 600, 900, 1800, 3600, 10800, 21600, 43200, 86400
	SuppressDuration int `json:"suppress_duration,omitempty"`
	// Specifies the alarm severity.
	// 1: critical, 2: major, 3: minor, 4: informational
	Level int `json:"level,omitempty"`
}

// Notification represents an alarm notification.
type Notification struct {
	// Specifies the notification type.
	// Possible values: notification, contact
	Type string `json:"type" required:"true"`
	// Specifies the list of SMN topic URNs.
	NotificationList []string `json:"notification_list" required:"true"`
}

// Create creates a new alarm rule.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v2/{project_id}/alarms
	raw, err := client.Post(client.ServiceURL("alarms"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		AlarmId string `json:"alarm_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.AlarmId, err
}
