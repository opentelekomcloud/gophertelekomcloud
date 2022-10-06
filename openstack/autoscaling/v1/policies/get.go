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
	ID             string         `json:"scaling_group_id"`
	Name           string         `json:"scaling_policy_name"`
	Status         string         `json:"policy_status"`
	Type           string         `json:"scaling_policy_type"`
	AlarmID        string         `json:"alarm_id"`
	SchedulePolicy SchedulePolicy `json:"scheduled_policy"`
	Action         Action         `json:"scaling_policy_action"`
	CoolDownTime   int            `json:"cool_down_time"`
	CreateTime     string         `json:"create_time"`
}

type SchedulePolicy struct {
	LaunchTime      string `json:"launch_time"`
	RecurrenceType  string `json:"recurrence_type"`
	RecurrenceValue string `json:"recurrence_value"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
}

type Action struct {
	Operation   string `json:"operation"`
	InstanceNum int    `json:"instance_number"`
}
