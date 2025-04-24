package alarm

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Keyword alarm rule ID.
	ID string `json:"keywords_alarm_rule_id" required:"true"`
	// Original rule name, which cannot be changed.
	Name string `json:"keywords_alarm_rule_name" required:"true"`
	// Rule name.
	Alias string `json:"alarm_rule_alias,omitempty"`
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

func UpdateKeywordRule(client *golangsdk.ServiceClient, opts UpdateOpts) (*KeywordRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/lts/alarms/keywords-alarm-rule
	raw, err := client.Put(client.ServiceURL("lts", "alarms", "keywords-alarm-rule"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"content-type": "application/json",
		},
	})
	if err != nil {
		return nil, err
	}

	var res KeywordRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type KeywordRule struct {
	// Keyword alarm rule ID.
	ID string `json:"keywords_alarm_rule_id"`
	// Original rule name, which cannot be changed.
	Name string `json:"keywords_alarm_rule_name"`
	// Rule name.
	Alias string `json:"alarm_rule_alias"`
	// Keyword alarm description.
	Description string `json:"keywords_alarm_rule_description"`
	// Keyword details.
	Details []Details `json:"keywords_requests"`
	// Alarm statistical period.
	Frequency *Frequency `json:"frequency"`
	// Alarm severity.
	Severity string `json:"keywords_alarm_level"`
	// Whether to send an alarm.
	Send bool `json:"keywords_alarm_send"`
	// Account ID
	DomainId string `json:"domain_id"`
	// Creation time (timestamp in milliseconds).
	CreatedAt int64 `json:"create_time"`
	// Update time (timestamp in milliseconds).
	UpdatedAt int64 `json:"update_time"`
	// Language of information added to emails.
	Language string `json:"language"`
	// Project ID.
	ProjectId string `json:"projectId"`
	// Notification topic
	Topics []Topics `json:"topics"`
	// Description
	ConditionExpression string `json:"condition_expression"`
	// Index ID.
	IndexId string `json:"indexId"`
	// Notification frequency, in minutes.
	NotificationFrequency int `json:"notification_frequency"`
	// Alarm action rule name.
	AlarmActionRuleName string `json:"alarm_action_rule_name"`
	// Message template name.
	TemplateName string `json:"template_name"`
	// Alarm status.
	Status string `json:"status"`
	// Number of queries in which the triggering condition is met. The default value is 1.
	TriggerConditionCount int `json:"trigger_condition_count"`
	// Number of times that log events meet the trigger condition. The default value is 1.
	TriggerConditionFrequency int `json:"trigger_condition_frequency"`
	// Whether to enable the alarm clearance notification. The default value is false.
	EnableRecoveryPolicy bool `json:"whether_recovery_policy"`
	// Number of queries in which the triggering condition is not met.
	// The alarm is cleared when this number reaches the value (3 by default) of this parameter.
	RecoveryPolicy int `json:"recovery_policy"`
}
