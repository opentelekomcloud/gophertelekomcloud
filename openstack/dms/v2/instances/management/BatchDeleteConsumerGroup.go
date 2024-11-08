package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchDeleteConsumberGroupOpts struct {
	// IDs of all consumer groups to be deleted.
	GroupIds []string `json:"group_ids" required:"true"`
}

// BatchDeleteConsumerGroup is used to delete multiple consumer groups of a Kafka instance in batches.
// Send POST /v2/{project_id}/instances/{instance_id}/groups/batch-delete
func BatchDeleteConsumerGroup(client *golangsdk.ServiceClient, instanceId string, opts BatchDeleteConsumberGroupOpts) (*BatchDeleteConsumerGroupResp, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("instances", instanceId, "groups", "batch-delete"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	if err != nil {
		return nil, err
	}

	var res BatchDeleteConsumerGroupResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchDeleteConsumerGroupResp struct {
	// List of consumer groups that failed to be deleted.
	FailedGroups []*FailedGroup `json:"failed_groups"`
	// Number of records that fail to be deleted.
	Total int `json:"total"`
}

type FailedGroup struct {
	// ID of consumer groups that failed to be deleted.
	GroupId string `json:"group_id"`
	// Cause of the deletion failure.
	ErrorMessage string `json:"error_message"`
}
