package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListConsumerGroupsOpts struct {
	// Offset, which is the position where the query starts. The value must be greater than or equal to 0.
	Offset string `q:"offset,omitempty"`
	// Querying partitions by top disk usage.
	Top string `q:"top,omitempty"`
	// Querying partitions by the percentage of the used disk space.
	Percentage string `q:"percentage,omitempty"`
}

// ListConsumerGroups is used to query all consumer groups.
// Send GET /v2/{project_id}/instances/{instance_id}/groups
func ListConsumerGroups(client *golangsdk.ServiceClient, instanceId string, opts ListConsumerGroupsOpts) (*ListConsumerGropusResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceId, "groups").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListConsumerGropusResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListConsumerGropusResp struct {
	// All consumer groups.
	Groups []*GroupInfo `json:"groups"`
	// Total number of consumer groups.
	Total int `json:"total"`
}

type GroupInfo struct {
	// Creation time.
	CreatedAt int64 `json:"created_at"`
	// Consumer group ID.
	GroupId string `json:"group_id"`
	// Consumer group status. The value can be:
	//    Dead: The consumer group has no members or metadata.
	//    Empty: The consumer group has metadata but has no members.
	//    PreparingRebalance: The consumer group is to be rebalanced.
	//    CompletingRebalance: All members have joined the group.
	//    Stable: Members in the consumer group can consume messages. "
	State string `json:"state"`
	// Coordinator ID.
	CoordinatorId int `json:"coordinator_id"`
	// Coordinator ID.
	Description string `json:"group_desc"`
	// Number of accumulated messages.
	Lag int64 `json:"lag"`
}
