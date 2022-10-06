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
	// Specifies the AS policy ID.
	PolicyID string `json:"scaling_policy_id"`
	// Specifies the AS policy name.
	// Supports fuzzy search.
	PolicyName string `json:"scaling_policy_name"`
	// Specifies the scaling resource ID.
	ResourceID string `json:"scaling_resource_id"`
	// Specifies the scaling resource type.
	// AS group: SCALING_GROUP
	// Bandwidth: BANDWIDTH
	ScalingResourceType string `json:"scaling_resource_type"`
	// Specifies the AS policy status.
	// INSERVICE: The AS policy is enabled.
	// PAUSED: The AS policy is disabled.
	// EXECUTING: The AS policy is being executed.
	PolicyStatus string `json:"policy_status"`
	// Specifies the AS policy type.
	// ALARM: indicates that the scaling action is triggered by an alarm.
	// A value is returned for alarm_id, and no value is returned for scheduled_policy.
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
	PolicyAction Action `json:"scaling_policy_action"`
	// Specifies the cooldown period (s).
	CoolDownTime int `json:"cool_down_time"`
	// Specifies the time when an AS policy was created. The time format complies with UTC.
	CreateTime string `json:"create_time"`
	// Provides additional information.
	Metadata Metadata `json:"meta_data"`
}

type SchedulePolicy struct {
	// Specifies the time when the scaling action is triggered. The time format complies with UTC.
	// If scaling_policy_type is set to SCHEDULED, the time format is YYYY-MM-DDThh:mmZ.
	// If scaling_policy_type is set to RECURRENCE, the time format is hh:mm.
	LaunchTime string `json:"launch_time"`
	// Specifies the type of a periodically triggered scaling action.
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
	// Specifies the start time of the scaling action triggered periodically. The time format complies with UTC.
	// The time format is YYYY-MM-DDThh:mmZ.
	StartTime string `json:"start_time"`
	// Specifies the end time of the scaling action triggered periodically. The time format complies with UTC.
	// The time format is YYYY-MM-DDThh:mmZ.
	EndTime string `json:"end_time"`
}

type Action struct {
	// Specifies the scaling action.
	// ADD: indicates adding instances.
	// REDUCE: indicates reducing instances.
	// SET: indicates setting the number of instances to a specified value.
	Operation string `json:"operation"`
	// Specifies the operation size.
	Size int `json:"size"`
	// Specifies the percentage of instances to be operated.
	Percentage int `json:"percentage"`
	// Specifies the operation restrictions.
	Limits int `json:"limits"`
}

type Metadata struct {
	// Specifies the bandwidth sharing type in the bandwidth scaling policy.
	BandwidthShareType string `json:"metadata_bandwidth_share_type"`
	// Specifies the EIP ID for the bandwidth in the bandwidth scaling policy.
	EipID string `json:"metadata_eip_id"`
	// Specifies the EIP for the bandwidth in the bandwidth scaling policy.
	EipAddress string `json:"metadata_eip_address"`
}
