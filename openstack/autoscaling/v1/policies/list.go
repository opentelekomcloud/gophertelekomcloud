package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Specifies the AS group ID.
	ScalingGroupId string
	// Specifies the AS policy name.
	// Supports fuzzy search.
	ScalingPolicyName string `q:"scaling_policy_name,omitempty"`
	// Specifies the AS policy type.
	// ALARM: alarm policy
	// SCHEDULED: scheduled policy
	// RECURRENCE: periodic policy
	ScalingPolicyType string `q:"scaling_policy_type,omitempty"`
	// Specifies the AS policy ID.
	ScalingPolicyId string `q:"scaling_policy_id,omitempty"`
	// Specifies the start line number. The default value is 0. The minimum parameter value is 0.
	StartNumber int32 `q:"start_number,omitempty"`
	// Specifies the number of query records. The default value is 20. The value range is 0 to 100.
	Limit int32 `q:"limit,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListScalingInstancesResponse, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /autoscaling-api/v1/{project_id}/scaling_policy/{scaling_group_id}/list
	raw, err := client.Get(client.ServiceURL("scaling_policy", opts.ScalingGroupId, "list")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListScalingInstancesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListScalingInstancesResponse struct {
	TotalNumber     int32                   `json:"total_number,omitempty"`
	StartNumber     int32                   `json:"start_number,omitempty"`
	Limit           int32                   `json:"limit,omitempty"`
	ScalingPolicies []ScalingV1PolicyDetail `json:"scaling_policies,omitempty"`
}

type ScalingV1PolicyDetail struct {
	// Specifies the AS group ID.
	ScalingGroupId string `json:"scaling_group_id,omitempty"`
	// Specifies the AS policy name.
	ScalingPolicyName string `json:"scaling_policy_name,omitempty"`
	// Specifies the AS policy ID.
	ScalingPolicyId string `json:"scaling_policy_id,omitempty"`
	// Specifies the AS policy status.
	// INSERVICE: The AS policy is enabled.
	// PAUSED: The AS policy is disabled.
	// EXECUTING: The AS policy is being executed.
	PolicyStatus string `json:"policy_status,omitempty"`
	// Specifies the AS policy type.
	// ALARM: indicates that the scaling action is triggered by an alarm.
	// A value is returned for alarm_id, and no value is returned for scheduled_policy.
	// SCHEDULED: indicates that the scaling action is triggered as scheduled. A value is returned for scheduled_policy,
	// and no value is returned for alarm_id, recurrence_type, recurrence_value, start_time, or end_time.
	// RECURRENCE: indicates that the scaling action is triggered periodically. Values are returned for scheduled_policy,
	// recurrence_type, recurrence_value, start_time, and end_time, and no value is returned for alarm_id.
	ScalingPolicyType string `json:"scaling_policy_type,omitempty"`
	// Specifies the alarm ID.
	AlarmId string `json:"alarm_id,omitempty"`
	// Specifies the periodic or scheduled AS policy.
	ScheduledPolicy ScheduledPolicy `json:"scheduled_policy,omitempty"`
	// Specifies the scaling action of the AS policy.
	ScalingPolicyAction ScalingPolicyActionV1 `json:"scaling_policy_action,omitempty"`
	// Specifies the cooldown period (s).
	CoolDownTime int32 `json:"cool_down_time,omitempty"`
	// Specifies the time when an AS policy was created. The time format complies with UTC.
	CreateTime string `json:"create_time,omitempty"`
}

type ScheduledPolicy struct {
	// Specifies the time when the scaling action is triggered. The time format complies with UTC.
	// If scaling_policy_type is set to SCHEDULED, the time format is YYYY-MM-DDThh:mmZ.
	// If scaling_policy_type is set to RECURRENCE, the time format is hh:mm.
	LaunchTime string `json:"launch_time"`
	// Specifies the type of periodically triggered scaling action.
	// Daily: indicates that the scaling action is triggered once a day.
	// Weekly: indicates that the scaling action is triggered once a week.
	// Monthly: indicates that the scaling action is triggered once a month.
	RecurrenceType string `json:"recurrence_type,omitempty"`
	// Specifies the frequency at which scaling actions are triggered.
	// If recurrence_type is set to Daily, the value is null, indicating that the scaling action is triggered once a day.
	// If recurrence_type is set to Weekly, the value ranges from 1 (Sunday) to 7 (Saturday).
	// The digits refer to dates in each week and separated by a comma, such as 1,3,5.
	// If recurrence_type is set to Monthly, the value ranges from 1 to 31.
	// The digits refer to the dates in each month and separated by a comma, such as 1,10,13,28.
	RecurrenceValue string `json:"recurrence_value,omitempty"`
	// Specifies the start time of the scaling action triggered periodically. The time format complies with UTC.
	// The time format is YYYY-MM-DDThh:mmZ.
	StartTime string `json:"start_time,omitempty"`
	// Specifies the end time of the scaling action triggered periodically. The time format complies with UTC.
	// The time format is YYYY-MM-DDThh:mmZ.
	EndTime string `json:"end_time,omitempty"`
}

type ScalingPolicyActionV1 struct {
	// Specifies the scaling action.
	// ADD: adds specified number of instances to the AS group.
	// REMOVE: removes specified number of instances from the AS group.
	// SET: sets the number of instances in the AS group.
	Operation string `json:"operation,omitempty"`
	// Specifies the number of instances to be operated.
	InstanceNumber int32 `json:"instance_number,omitempty"`
	// Specifies the percentage of instances to be operated.
	InstancePercentage int32 `json:"instance_percentage,omitempty"`
}
