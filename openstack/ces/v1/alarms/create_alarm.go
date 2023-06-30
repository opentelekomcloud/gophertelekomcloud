package alarms

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateAlarmOpts struct {
	// Specifies the alarm rule name.
	// Enter 1 to 128 characters. Only letters, digits, underscores (_), and hyphens (-) are allowed.
	AlarmName string `json:"alarm_name"`
	// Provides supplementary information about the alarm rule. Enter 0 to 256 characters.
	AlarmDescription string `json:"alarm_description,omitempty"`
	// Specifies the alarm metric.
	Metric MetricForAlarm `json:"metric"`
	// Specifies the alarm triggering condition.
	Condition Condition `json:"condition"`
	// Specifies whether to enable the alarm.
	AlarmEnabled *bool `json:"alarm_enabled,omitempty"`
	// Specifies whether to enable the action to be triggered by an alarm. The default value is true.
	AlarmActionEnabled *bool `json:"alarm_action_enabled,omitempty"`
	// Specifies the alarm severity. Possible values are 1, 2 (default), 3 and 4,
	// indicating critical, major, minor, and informational, respectively.
	AlarmLevel int `json:"alarm_level,omitempty"`
	// Specifies the action to be triggered by an alarm.
	AlarmType string `json:"alarm_type,omitempty"`
	// Specifies the action to be triggered by an alarm.
	AlarmActions []AlarmActions `json:"alarm_actions,omitempty"`
	// Specifies the action to be triggered after the alarm is cleared.
	OkActions []AlarmActions `json:"ok_actions,omitempty"`
}

type MetricForAlarm struct {
	// Specifies the namespace of a service.
	// The value must be in the service.item format and can contain 3 to 32 characters.
	// service and item each must start with a letter and contain only letters, digits, and underscores (_).
	Namespace string `json:"namespace"`
	// Specifies the metric name.
	// Start with a letter. Enter 1 to 64 characters. Only letters, digits, and underscores (_) are allowed.
	MetricName string `json:"metric_name"`
	// Specifies the list of metric dimensions.
	Dimensions []MetricsDimension `json:"dimensions,omitempty"`
	// Specifies the resource group ID selected during the alarm rule creation.
	ResourceGroupId string `json:"resource_group_id,omitempty"`
}

func CreateAlarm(client *golangsdk.ServiceClient, opts CreateAlarmOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /V1.0/{project_id}/alarms
	raw, err := client.Post(client.ServiceURL("alarms"), b, nil, nil)
	if err != nil {
		return "", err
	}

	var s struct {
		AlarmId string `json:"alarm_id"`
	}
	err = extract.Into(raw.Body, &s)
	return s.AlarmId, err
}
