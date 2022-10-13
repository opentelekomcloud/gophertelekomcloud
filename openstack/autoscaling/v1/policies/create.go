package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the AS policy name. The name contains only letters, digits, underscores (_), and hyphens (-), and cannot exceed 64 characters.
	Name string `json:"scaling_policy_name" required:"true"`
	// Specifies the AS group ID, which can be obtained using the API for querying AS groups.
	ID string `json:"scaling_group_id" required:"true"`
	// Specifies the AS policy type.
	// ALARM (corresponding to alarm_id): indicates that the scaling action is triggered by an alarm.
	// SCHEDULED (corresponding to scheduled_policy): indicates that the scaling action is triggered as scheduled.
	// RECURRENCE (corresponding to scheduled_policy): indicates that the scaling action is triggered periodically.
	Type string `json:"scaling_policy_type" required:"true"`
	// Specifies the alarm rule ID. This parameter is mandatory when scaling_policy_type is set to ALARM.
	// After this parameter is specified, the value of scheduled_policy does not take effect.
	// After you create an alarm policy, the system automatically adds an alarm triggering
	// activity of the autoscaling type to the alarm_actions field in the alarm rule specified by the parameter value.
	AlarmID string `json:"alarm_id,omitempty"`
	// Specifies the periodic or scheduled AS policy. This parameter is mandatory
	// when scaling_policy_type is set to SCHEDULED or RECURRENCE.
	// After this parameter is specified, the value of alarm_id does not take effect.
	SchedulePolicy SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	// Specifies the scaling action of the AS policy.
	Action Action `json:"scaling_policy_action,omitempty"`
	// Specifies the cooldown period (in seconds). The value ranges from 0 to 86400 and is 300 by default.
	CoolDownTime int `json:"cool_down_time,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	raw, err := client.Post(client.ServiceURL("scaling_policy"), b, nil, &golangsdk.RequestOpts{
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
