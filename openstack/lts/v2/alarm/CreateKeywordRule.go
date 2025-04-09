package alarm

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Keyword alarm rule names.
	Name string `json:"keywords_alarm_rule_name" required:"true"`
	// Keyword alarm description.
	Description string `json:"keywords_alarm_rule_description,omitempty"`
	// Keyword details.
	Details []Details `json:"keywords_requests" required:"true"`
	// Alarm statistical period.
	Frequency *Frequency `json:"frequency" required:"true"`
	// Alarm severity.
	Severity string `json:"keywords_alarm_level" required:"true"`
	// Whether to send an alarm.
	Send *bool `json:"keywords_alarm_send" required:"true"`
	// Account ID
	DomainId string `json:"domain_id" required:"true"`
	// Number of times that log events meet the trigger condition. The default value is 1.
	TriggerConditionCount int `json:"trigger_condition_count,omitempty"`
	// Number of queries in which the triggering condition is met. The default value is 1.
	TriggerConditionFrequency int `json:"trigger_condition_frequency,omitempty"`
	// Whether to enable the alarm clearance notification. The default value is false.
	EnableRecoveryPolicy *bool `json:"whether_recovery_policy,omitempty"`
	// Number of queries in which the triggering condition is not met.
	// The alarm is cleared when this number reaches the value (3 by default) of this parameter.
	RecoveryPolicy int `json:"recovery_policy,omitempty"`
	// Notification frequency, in minutes.
	NotificationFrequency int `json:"notification_frequency" required:"true"`
	// Alarm action rule name.
	AlarmActionRuleName string `json:"alarm_action_rule_name,omitempty"`
	// Notification topic
	NotificationSave *NotificationSave `json:"notification_save_rule,omitempty"`
}

type Details struct {
	// Log stream ID.
	LogStreamId string `json:"log_stream_id" required:"true"`
	// Log stream name.
	LogStreamName string `json:"log_stream_name,omitempty"`
	// Log group ID.
	LogGroupId string `json:"log_group_id" required:"true"`
	// Log group name.
	LogGroupName string `json:"log_group_name,omitempty"`
	// Keyword.
	Keyword string `json:"keywords" required:"true"`
	// Condition.
	Condition string `json:"condition" required:"true"`
	// Keyword threshold, which forms a condition with keyword and condition. An alarm is triggered when the condition is met.
	Number int `json:"number" required:"true"`
	// Time range for querying the latest data when a task is executed.
	SearchTimeRange int `json:"search_time_range" required:"true"`
	// Query time unit.
	SearchTimeRangeUnit string `json:"search_time_range_unit" required:"true"`
}

type Frequency struct {
	Type string `json:"type" required:"true"`
	// Cron expression, which uses the 24-hour format and is precise down to the minute.
	CronExpr string `json:"cron_expr,omitempty"`
	// This field is used when type is set to DAILY or WEEKLY.
	// ranges from 0 to 23.
	HourOfDay int `json:"hour_of_day,omitempty"`
	// This field is used when type is set to WEEKLY (from Sunday to Saturday).
	DayOfWeek int `json:"day_of_week,omitempty"`
	// Value of a period. This field is used when type is set to FIXED_RATE.
	// It is used together with fixed_rate_unit to indicate a fixed period.
	FixedRate int `json:"fixed_rate,omitempty"`
	// Unit of a period. This field is used when type is set to FIXED_RATE.
	// It is used together with fixed_rate to indicate a fixed period.
	FixedRateUnit string `json:"fixed_rate_unit,omitempty"`
}

type NotificationSave struct {
	// Language of the preference.
	Language string `json:"language" required:"true"`
	// Time zone information used in a notification. Example: +08:00
	Timezone string `json:"timezone,omitempty"`
	// Username used in a notification. It is generally displayed in the first line of the greeting.
	UserName string `json:"user_name" required:"true"`
	// Topic information
	Topics []TopicsCreate `json:"topics" required:"true"`
	// Message template name.
	TemplateName string `json:"template_name,omitempty"`
}

type TopicsCreate struct {
	// Topic name.
	Name string `json:"name" required:"true"`
	// Specifies the resource identifier of the topic, which is unique.
	TopicUrn string `json:"topic_urn" required:"true"`
	// Specifies the topic display name, which is presented as the name of the email sender in email messages.
	DisplayName string `json:"display_name,omitempty"`
	// Specifies the message push policy.
	PushPolicy int `json:"push_policy,omitempty"`
}

func CreateKeywordRule(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST /v2/{project_id}/lts/alarms/keywords-alarm-rule
	raw, err := client.Post(client.ServiceURL("lts", "alarms", "keywords-alarm-rule"), b, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"keywords_alarm_rule_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
