package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetConsumerGroupDetails is used to query consumer group details.
// Send GET /v2/{project_id}/instances/{instance_id}/management/groups/{group}
func GetConsumerGroupDetails(client *golangsdk.ServiceClient, instanceId, groupId string) (*ConsumerGroupResp, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "management", "groups", groupId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ConsumerGroupResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ConsumerGroupResp struct {
	// Details of the result of the cross-VPC access modification.
	Group *Group `json:"group"`
}

type Group struct {
	// Consumer group name.
	GroupId string `json:"group_id"`
	// Consumer group status. The value can be:
	//    Dead: The consumer group has no members and no metadata.
	//    Empty: The consumer group has metadata but has no members.
	//    PreparingRebalance: The consumer group is to be rebalanced.
	//    CompletingRebalance: All members have jointed the group.
	//    Stable: Members in the consumer group can consume messages normally.
	State string `json:"state"`
	// Coordinator ID.
	CoordinatorId int `json:"coordinator_id"`
	// Consumer list.
	Members []*Member `json:"members"`
	// Consumer offset.
	GroupMessageOffsets []GroupMessageOffest `json:"group_message_offsets"`
	// Partition assignment policy.
	AssignmentStrategy string `json:"assignment_strategy"`
}

type Member struct {
	// Consumer address.
	Host string `json:"host"`
	// Details about the partition assigned to the consumer.
	Assignment []*Assignment `json:"assignment"`
	// Consumer ID.
	MemberId string `json:"member_id"`
	// Client ID.
	ClientId string `json:"client_id"`
}
type Assignment struct {
	// Topic name.
	Topic string `json:"topic"`
	// Partition list.
	Partitions []int `json:"partitions"`
}

type GroupMessageOffest struct {
	// Partition number.
	Partition int `json:"partition"`
	// Number of remaining messages that can be retrieved, that is, the number of accumulated messages.
	Lag int64 `json:"lag"`
	// Topic name.
	Topic string `json:"topic"`
	// Consumer offset.
	MessageCurrentOffset int64 `json:"message_current_offset"`
	// Log end offset (LEO).
	MessageLogEndOffset int64 `json:"message_log_end_offset"`
}
