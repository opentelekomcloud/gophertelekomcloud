package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type SchedulePolicyOpts struct {
	LaunchTime      string `json:"launch_time" required:"true"`
	RecurrenceType  string `json:"recurrence_type,omitempty"`
	RecurrenceValue string `json:"recurrence_value,omitempty"`
	StartTime       string `json:"start_time,omitempty"`
	EndTime         string `json:"end_time,omitempty"`
}

type ActionOpts struct {
	Operation   string `json:"operation,omitempty"`
	InstanceNum int    `json:"instance_number,omitempty"`
}

type UpdateOpts struct {
	Name           string             `json:"scaling_policy_name,omitempty"`
	Type           string             `json:"scaling_policy_type,omitempty"`
	AlarmID        string             `json:"alarm_id,omitempty"`
	SchedulePolicy SchedulePolicyOpts `json:"scheduled_policy,omitempty"`
	Action         ActionOpts         `json:"scaling_policy_action,omitempty"`
	CoolDownTime   int                `json:"cool_down_time,omitempty"`
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (string, error) {
	body, err := golangsdk.BuildRequestBody(opts, "")
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
