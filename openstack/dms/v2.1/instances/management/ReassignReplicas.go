package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

const (
	reassignmentPath = "replicas-reassignment"
	topicPath        = "topics"
)

type ReassignReplicasOpts struct {
	// Assignment of replicas of the partition after the change.
	Partitions []*Partition `json:"partitions"`
}

type Partition struct {
	// Partition ID.
	PartitionID int `json:"partition_id"`
	// ID of the broker where the replica is expected to reside. The first integer in the array represents the leader replica broker ID. All partitions must have the same number of replicas. The number of replicas cannot be larger than the number of brokers.
	Replicas []int `json:"replicas"`
}

// ReassignReplicas is used to reassign replicas of a topic for a Kafka instance.
// Send POST /v2/{project_id}/instances/{instance_id}/management/topics/{topic}/replicas-reassignment
func ReassignReplicas(client *golangsdk.ServiceClient, instanceId, topicId string, opts ResetMessageOffsetOpts) error {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL(instances.ResourcePath, instanceId, managementPath, topicPath, topicId, reassignmentPath), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})

	return err
}
