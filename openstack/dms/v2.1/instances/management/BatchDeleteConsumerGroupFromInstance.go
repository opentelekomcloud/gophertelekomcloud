package management

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/dms/v2.1/instances"
)

const batchDeletePath = "batch-delete"

type BatchDeleteConsumberGroupOpts struct {
	// IDs of all consumer groups to be deleted.
	GroupIds []string `json:"group_ids" required:"true"`
}

// BatchDeleteConsumerGroupFromInstance is used to delete multiple consumer groups of a Kafka instance in batches.
// Send POST /v2/{project_id}/instances/{instance_id}/groups/batch-delete
func BatchDeleteConsumerGroupFromInstance(client *golangsdk.ServiceClient, instanceId, groupId string, opts BatchDeleteConsumberGroupOpts) (*BatchDeleteConsumerGroupResp, error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(instances.ResourcePath, instanceId, groupPath, batchDeletePath), body, nil, &golangsdk.RequestOpts{
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
