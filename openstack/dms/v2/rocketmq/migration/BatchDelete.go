package migration

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchDeleteOpts struct {
	// IDs of the tasks to be deleted.
	TaskIDs []string `json:"task_ids" required:"true"`
}

// BatchDeleteResponse is returned by BatchDelete.
type BatchDeleteResponse struct {
	// IDs of the tasks that are successfully deleted.
	SuccessTaskList []string `json:"success_task_list"`
}

// BatchDelete deletes metadata migration tasks of a RocketMQ instance in batches.
// Send POST to /v2/{project_id}/instances/{instance_id}/metadata/batch-delete
func BatchDelete(client *golangsdk.ServiceClient, instanceID string, opts BatchDeleteOpts) (*BatchDeleteResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceID, "metadata", "batch-delete").Build()
	if err != nil {
		return nil, err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res BatchDeleteResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
