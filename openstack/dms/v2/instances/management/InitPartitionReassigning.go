package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type InitPartitionReassigningOpts struct {
	// Reassignment plan.
	Reassignments []PartitionReassign `json:"reassignments" required:"true"`
	// Reassignment threshold.
	Throttle int `json:"throttle,omitempty"`
	// Indicates whether the task is scheduled. If no, is_schedule and execute_at can be left blank. If yes, is_schedule is true and execute_at must be specified.
	IsSchedule bool `json:"is_schedule,omitempty"`
	// Schedule time. The value is a UNIX timestamp, in ms.
	ExecuteAt int64 `json:"execute_at,omitempty"`
	// Set true to perform time estimation tasks and false to perform rebalancing tasks.
	// Default: false
	TimeEstimate bool `json:"time_estimate,omitempty"`
}

type PartitionReassign struct {
	// Topic name.
	Topic string `json:"topic" required:"true"`
	// List of brokers to which partitions are reassigned. This parameter is mandatory in automatic assignment.
	Brokers []int `json:"brokers,omitempty"`
	// Replication factor, which can be specified in automatic assignment.
	ReplicationFactor int `json:"replication_factor,omitempty"`
	// Manually specified assignment plan. The brokers parameter and this parameter cannot be empty at the same time.
	Assignment []*TopicAssignment `json:"assignment,omitempty"`
}

type TopicAssignment struct {
	// Partition number in manual assignment.
	Partition int `json:"partition,omitempty"`
	// List of brokers to be assigned to a partition in manual assignment.
	PartitionBrokers []int `json:"partition_brokers,omitempty"`
}

// InitPartitionReassigning is used to submit a partition rebalancing task to a Kafka instance or calculate estimated rebalancing time.
// Send POST /v2/kafka/{project_id}/instances/{instance_id}/reassign
func InitPartitionReassigning(client *golangsdk.ServiceClient, instanceId string, opts *InitPartitionReassigningOpts) (*InitResp, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", instanceId, "reassign"), body, nil, &golangsdk.RequestOpts{})
	if err != nil {
		return nil, err
	}

	var res InitResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type InitResp struct {
	// Task ID. Only job_id is returned for a rebalancing task.
	JobId string `json:"job_id"`
	// Estimated time, in seconds. Only reassignment_time is returned for a time estimation task.
	ReassignmentTime int `json:"reassignment_time"`
}
