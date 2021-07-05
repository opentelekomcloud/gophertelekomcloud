package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateResult is a struct which represents the create result of policy
type CreateResult struct {
	golangsdk.Result
}

// Extract of CreateResult will deserialize the result of Creation
func (r CreateResult) Extract() (string, error) {
	var a struct {
		ID string `json:"scaling_policy_id"`
	}
	err := r.Result.ExtractInto(&a)
	return a.ID, err
}

// DeleteResult is a struct which represents the delete result.
type DeleteResult struct {
	golangsdk.ErrResult
}

// Policy is a struct that represents the result of get policy
type Policy struct {
	PolicyID            string         `json:"scaling_policy_id"`
	PolicyName          string         `json:"scaling_policy_name"`
	ResourceID          string         `json:"scaling_resource_id"`
	ScalingResourceType string         `json:"scaling_resource_type"`
	PolicyStatus        string         `json:"policy_status"`
	Type                string         `json:"scaling_policy_type"`
	AlarmID             string         `json:"alarm_id"`
	SchedulePolicy      SchedulePolicy `json:"scheduled_policy"`
	PolicyAction        Action         `json:"scaling_policy_action"`
	CoolDownTime        int            `json:"cool_down_time"`
	CreateTime          string         `json:"create_time"`
	Metadata            Metadata       `json:"meta_data"`
}

type SchedulePolicy struct {
	LaunchTime      string `json:"launch_time"`
	RecurrenceType  string `json:"recurrence_type"`
	RecurrenceValue string `json:"recurrence_value"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
}

type Action struct {
	Operation  string `json:"operation"`
	Size       int    `json:"size"`
	Percentage int    `json:"percentage"`
	Limits     int    `json:"limits"`
}

type Metadata struct {
	BandwidthShareType string `json:"metadata_bandwidth_share_type"`
	EipID              string `json:"metadata_eip_id"`
	EipAddress         string `json:"metadata_eip_address"`
}

// GetResult is a struct which represents the get result
type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (Policy, error) {
	var p Policy
	err := r.Result.ExtractIntoStructPtr(&p, "scaling_policy")
	return p, err
}

// UpdateResult is a struct from which can get the result of update method
type UpdateResult struct {
	golangsdk.Result
}

// Extract will deserialize the result to group id with string
func (r UpdateResult) Extract() (string, error) {
	var a struct {
		ID string `json:"scaling_policy_id"`
	}
	err := r.Result.ExtractInto(&a)
	return a.ID, err
}
