package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Specifies the AS policy name. The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	Name string `json:"scaling_policy_name,omitempty"`
	// Specifies the AS policy type.
	// ALARM (corresponding to alarm_id): indicates that the scaling action is triggered by an alarm.
	// SCHEDULED (corresponding to scheduled_policy): indicates that the scaling action is triggered as scheduled.
	// RECURRENCE (corresponding to scheduled_policy): indicates that the scaling action is triggered periodically.
	Type string `json:"scaling_policy_type,omitempty"`
	// Specifies the alarm rule ID. This parameter is mandatory when scaling_policy_type is set to ALARM.
	// After this parameter is specified, the value of scheduled_policy does not take effect.
	// After you modify an alarm policy, the system automatically adds an alarm triggering activity of the autoscaling
	// type to the alarm_actions field in the alarm rule specified by the parameter value.
	// You can obtain the parameter value by querying Cloud Eye alarm rules.
	AlarmID string `json:"alarm_id,omitempty"`
	// Specifies the periodic or scheduled AS policy. This parameter is mandatory when scaling_policy_type is set to SCHEDULED or RECURRENCE.
	// After this parameter is specified, the value of alarm_id does not take effect.
	SchedulePolicy SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	// Specifies the scaling action of the AS policy.
	Action Action `json:"scaling_policy_action,omitempty"`
	// Specifies the cooldown period (in seconds). The value ranges from 0 to 86400.
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

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (string, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Put(client.ServiceURL("scaling_policy", id), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return "", err
	}

	var res struct {
		ID string `json:"scaling_policy_id"`
	}
	err = extract.Into(raw.Body, &res)
	return res.ID, err
}
