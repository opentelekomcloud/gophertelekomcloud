package policies

type PolicyOpts struct {
	// Specifies the AS policy name. The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	PolicyName string `json:"scaling_policy_name"`
	// Specifies the AS policy type.
	// ALARM (corresponding to alarm_id): indicates that the scaling action is triggered by an alarm.
	// SCHEDULED (corresponding to scheduled_policy): indicates that the scaling action is triggered as scheduled.
	// RECURRENCE (corresponding to scheduled_policy): indicates that the scaling action is triggered periodically.
	PolicyType string `json:"scaling_policy_type"`
	// Specifies the scaling resource ID, which is the unique ID of an AS group or bandwidth.
	// If scaling_resource_type is set to SCALING_GROUP, this parameter indicates the unique AS group ID.
	// If scaling_resource_type is set to BANDWIDTH, this parameter indicates the unique bandwidth ID.
	ResourceID string `json:"scaling_resource_id"`
	// Specifies the scaling resource type.
	// AS group: SCALING_GROUP
	// Bandwidth: BANDWIDTH
	ResourceType string `json:"scaling_resource_type"`
	// Specifies the alarm rule ID. This parameter is mandatory when scaling_policy_type is set to ALARM.
	// After this parameter is specified, the value of scheduled_policy does not take effect.
	// After you create an alarm policy, the system automatically adds an alarm triggering
	// activity of the autoscaling type to the alarm_actions field in the alarm rule specified by the parameter value.
	// You can obtain the parameter value by querying Cloud Eye alarm rules.
	AlarmID string `json:"alarm_id,omitempty"`
	// Specifies the periodic or scheduled AS policy.
	// This parameter is mandatory when scaling_policy_type is set to SCHEDULED or RECURRENCE.
	// After this parameter is specified, the value of alarm_id does not take effect.
	SchedulePolicy *SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	// Specifies the scaling action of the AS policy.
	PolicyAction *ActionOpts `json:"scaling_policy_action,omitempty"`
	// Specifies the cooldown period (in seconds). The value ranges from 0 to 86400 and is 300 by default.
	CoolDownTime int `json:"cool_down_time,omitempty"`
}

type SchedulePolicyOpts struct {
	// Specifies the time when the scaling action is triggered. The time format complies with UTC.
	// If scaling_policy_type is set to SCHEDULED, the time format is YYYY-MM-DDThh:mmZ.
	// If scaling_policy_type is set to RECURRENCE, the time format is hh:mm.
	LaunchTime string `json:"launch_time" required:"true"`
	// Specifies the periodic triggering type. This parameter is mandatory when scaling_policy_type is set to RECURRENCE.
	// Daily: indicates that the scaling action is triggered once a day.
	// Weekly: indicates that the scaling action is triggered once a week.
	// Monthly: indicates that the scaling action is triggered once a month.
	RecurrenceType string `json:"recurrence_type,omitempty"`
	// Specifies the day when a periodic scaling action is triggered. This parameter is mandatory when scaling_policy_type is set to RECURRENCE.
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
	// This parameter is mandatory when scaling_policy_type is set to RECURRENCE.
	// When the scaling action is triggered periodically, the end time cannot be earlier than the current and start time.
	// The time format is YYYY-MM-DDThh:mmZ.
	EndTime string `json:"end_time,omitempty"`
}

type ActionOpts struct {
	// Specifies the operation to be performed. The default operation is ADD.
	// If scaling_resource_type is set to SCALING_GROUP, the following operations are supported:
	// ADD: indicates adding instances.
	// REMOVE/REDUCE: indicates removing or reducing instances.
	// SET: indicates setting the number of instances to a specified value.
	// If scaling_resource_type is set to BANDWIDTH, the following operations are supported:
	// ADD: indicates adding instances.
	// REDUCE: indicates reducing instances.
	// SET: indicates setting the number of instances to a specified value.
	Operation string `json:"operation,omitempty"`
	// Specifies the operation size. The value is an integer from 0 to 300. The default value is 1.
	// This parameter can be set to 0 only when operation is set to SET.
	// If scaling_resource_type is set to SCALING_GROUP, this parameter indicates the number of instances.
	// The value is an integer from 0 to 300 and the default value is 1.
	// If scaling_resource_type is set to BANDWIDTH, this parameter indicates the bandwidth (Mbit/s).
	// The value is an integer from 1 to 300 and the default value is 1.
	// If scaling_resource_type is set to SCALING_GROUP, either size or percentage can be set.
	Size int `json:"size,omitempty"`
	// Specifies the percentage of instances to be operated. If operation is set to ADD, REMOVE, or REDUCE,
	// the value of this parameter is an integer from 1 to 20000. If operation is set to SET, the value is an integer from 0 to 20000.
	// If scaling_resource_type is set to SCALING_GROUP, either size or percentage can be set.
	// If neither size nor percentage is set, the default value of size is 1.
	// If scaling_resource_type is set to BANDWIDTH, percentage is unavailable.
	Percentage int `json:"percentage,omitempty"`
	// Specifies the operation restrictions.
	// If scaling_resource_type is set to BANDWIDTH and operation is not SET, this parameter takes effect and the unit is Mbit/s.
	// If operation is set to ADD, this parameter indicates the maximum bandwidth allowed.
	// If operation is set to REDUCE, this parameter indicates the minimum bandwidth allowed.
	Limits int `json:"limits,omitempty"`
}
