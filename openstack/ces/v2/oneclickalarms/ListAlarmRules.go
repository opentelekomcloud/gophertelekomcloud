package oneclickalarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// AlarmRule represents an alarm rule in one-click monitoring.
type AlarmRule struct {
	// Specifies the alarm rule ID.
	AlarmId string `json:"alarm_id"`
	// Specifies the alarm rule name.
	Name string `json:"name"`
	// Provides supplementary information about the alarm rule.
	Description string `json:"description"`
	// Specifies the namespace of a service.
	Namespace string `json:"namespace"`
	// Specifies the alarm policies.
	Policies []Policy `json:"policies"`
	// Specifies the resource list.
	Resources []Resource `json:"resources"`
	// Specifies the alarm rule type.
	// Possible values: ALL_INSTANCE, RESOURCE_GROUP, MULTI_INSTANCE, EVENT.SYS, EVENT.CUSTOM
	Type string `json:"type"`
	// Specifies whether to generate alarms.
	Enabled bool `json:"enabled"`
	// Specifies whether notification is enabled.
	NotificationEnabled bool `json:"notification_enabled"`
	// Specifies the action triggered by an alarm.
	AlarmNotifications []Notification `json:"alarm_notifications"`
	// Specifies the action triggered after an alarm is cleared.
	OkNotifications []Notification `json:"ok_notifications"`
	// Specifies the time when the alarm notification was enabled.
	// The value is in the format of HH:MM.
	NotificationBeginTime string `json:"notification_begin_time"`
	// Specifies the time when the alarm notification was disabled.
	// The value is in the format of HH:MM.
	NotificationEndTime string `json:"notification_end_time"`
}

// Policy represents an alarm policy.
type Policy struct {
	// Specifies the alarm policy ID.
	AlarmPolicyId string `json:"alarm_policy_id"`
	// Specifies the metric name.
	MetricName string `json:"metric_name"`
	// Specifies the rollup period in seconds.
	// Possible values: 0, 1, 300, 1200, 3600, 14400, 86400
	Period int `json:"period"`
	// Specifies the data rollup method.
	// Possible values: average, max, min, sum, variance
	Filter string `json:"filter"`
	// Specifies the comparison operator.
	// Possible values: >, <, >=, <=, =, !=
	ComparisonOperator string `json:"comparison_operator"`
	// Specifies the alarm threshold.
	Value float64 `json:"value"`
	// Specifies the data unit.
	Unit string `json:"unit"`
	// Specifies the number of consecutive times that the alarm condition is met.
	Count int `json:"count"`
	// Specifies the interval for triggering an alarm if the alarm persists in seconds.
	// Possible values: 0, 300, 600, 900, 1800, 3600, 10800, 21600, 43200, 86400
	SuppressDuration int `json:"suppress_duration"`
	// Specifies the alarm severity.
	// 1: critical, 2: major, 3: minor, 4: informational
	Level int `json:"level"`
	// Specifies whether the alarm policy is enabled.
	Enabled bool `json:"enabled"`
}

// Resource represents a resource in the alarm rule.
type Resource struct {
	// Specifies the resource group ID.
	ResourceGroupId string `json:"resource_group_id"`
	// Specifies the resource group name.
	ResourceGroupName string `json:"resource_group_name"`
	// Specifies the metric dimensions.
	Dimensions []Dimension `json:"dimensions"`
}

// Dimension represents a metric dimension.
type Dimension struct {
	// Specifies the dimension name.
	Name string `json:"name"`
	// Specifies the dimension value.
	Value string `json:"value"`
}

// ListAlarmRules queries the alarm rules of one service in one-click monitoring.
func ListAlarmRules(client *golangsdk.ServiceClient, oneClickAlarmId string) ([]AlarmRule, error) {
	// GET /v2/{project_id}/one-click-alarms/{one_click_alarm_id}/alarms
	raw, err := client.Get(client.ServiceURL("one-click-alarms", oneClickAlarmId, "alarms"), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		Alarms []AlarmRule `json:"alarms"`
	}
	err = extract.Into(raw.Body, &res)
	return res.Alarms, err
}
