package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Policy, error) {
	raw, err := client.Get(client.ServiceURL("scaling_policy", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Policy
	err = extract.IntoStructPtr(raw.Body, &res, "scaling_policy")
	return &res, err
}

type Policy struct {
	// Specifies the AS group ID.
	ID string `json:"scaling_group_id"`
	// Specifies the AS policy name.
	Name string `json:"scaling_policy_name"`
	// Specifies the AS policy status.
	// INSERVICE: The AS policy is enabled.
	// PAUSED: The AS policy is disabled.
	// EXECUTING: The AS policy is being executed.
	Status string `json:"policy_status"`
	// Specifies the AS policy type.
	// ALARM: indicates that the scaling action is triggered by an alarm. A value is returned for alarm_id, and no value is returned for scheduled_policy.
	// SCHEDULED: indicates that the scaling action is triggered as scheduled.
	// A value is returned for scheduled_policy, and no value is returned for alarm_id, recurrence_type, recurrence_value, start_time, or end_time.
	// RECURRENCE: indicates that the scaling action is triggered periodically.
	// Values are returned for scheduled_policy, recurrence_type, recurrence_value, start_time, and end_time, and no value is returned for alarm_id.
	Type string `json:"scaling_policy_type"`
	// Specifies the alarm ID.
	AlarmID string `json:"alarm_id"`
	// Specifies the periodic or scheduled AS policy.
	SchedulePolicy SchedulePolicy `json:"scheduled_policy"`
	// Specifies the scaling action of the AS policy.
	Action Action `json:"scaling_policy_action"`
	// Specifies the cooldown period (s).
	CoolDownTime int `json:"cool_down_time"`
	// Specifies the time when an AS policy was created. The time format complies with UTC.
	CreateTime string `json:"create_time"`
}

type SchedulePolicy struct {
	// Specifies the time when the scaling action is triggered. The time format complies with UTC.
	// If scaling_policy_type is set to SCHEDULED, the time format is YYYY-MM-DDThh:mmZ.
	// If scaling_policy_type is set to RECURRENCE, the time format is hh:mm.
	LaunchTime string `json:"launch_time"`
	// Specifies the type of periodically triggered scaling action.
	// Daily: indicates that the scaling action is triggered once a day.
	// Weekly: indicates that the scaling action is triggered once a week.
	// Monthly: indicates that the scaling action is triggered once a month.
	RecurrenceType string `json:"recurrence_type"`
	// Specifies the frequency at which scaling actions are triggered.
	// If recurrence_type is set to Daily, the value is null, indicating that the scaling action is triggered once a day.
	// If recurrence_type is set to Weekly, the value ranges from 1 (Sunday) to 7 (Saturday).
	// The digits refer to dates in each week and separated by a comma, such as 1,3,5.
	// If recurrence_type is set to Monthly, the value ranges from 1 to 31.
	// The digits refer to the dates in each month and separated by a comma, such as 1,10,13,28.
	RecurrenceValue string `json:"recurrence_value"`
	// Specifies the start time of the scaling action triggered periodically.
	// The time format complies with UTC.
	// The time format is YYYY-MM-DDThh:mmZ.
	StartTime string `json:"start_time"`
	// Specifies the end time of the scaling action triggered periodically.
	// The time format complies with UTC.
	// The time format is YYYY-MM-DDThh:mmZ.
	EndTime string `json:"end_time"`
}

type Action struct {
	// Specifies the operation to be performed. The default operation is ADD.
	// ADD: adds specified number of instances to the AS group.
	// REMOVE/REDUCE: removes or reduces specified number of instances from the AS group.
	// SET: sets the number of instances in the AS group.
	Operation string `json:"operation,omitempty"`
	// Specifies the number of instances to be operated. The default number is 1.
	// The value range is as follows for a default quota:
	// If operation is set to SET, the value range is 0 to 200.
	// If operation is set to ADD, REMOVE, or REDUCE, the value range is 1 to 200.
	// NOTE:
	// Either instance_number or instance_percentage is required.
	InstanceNum int `json:"instance_number,omitempty"`
	// Specifies the percentage of instances to be operated. You can increase, decrease,
	// or set the number of instances in an AS group to the specified percentage of the current number of instances.
	// If operation is set to ADD, REMOVE or REDUCE, the value of this parameter is an integer from 1 to 20000.
	// If operation is set to SET, the value is an integer from 0 to 20000.
	// If neither instance_number nor instance_percentage is specified, the number of instances to be operated is 1.
	// Either instance_number or instance_percentage is required.
	InstancePercentage int `json:"instance_percentage,omitempty"`
}
