package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchDeleteTasksOpts struct {
	Jobs []DeleteJobReq `json:"jobs"`
}

type DeleteJobReq struct {
	// The value can be terminated, force_terminate, or delete. terminate indicates that the migration task is stopped,
	// force_terminate indicates that the migration task is forcibly stopped, and delete indicates that the migration task is deleted.
	// Values: terminate force_terminate delete
	DeleteType string `json:"delete_type"`
	// Task ID.
	JobId string `json:"job_id"`
}

type BatchDeleteJobsResponse struct {
	Results []IdJobResp `json:"results,omitempty"`
	Count   int         `json:"count,omitempty"`
}

func BatchDeleteTasks(client *golangsdk.ServiceClient, opts BatchDeleteTasksOpts) (*BatchTasksResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// DELETE /v3/{project_id}/jobs/batch-jobs
	raw, err := client.DeleteWithBody(client.ServiceURL("jobs", "batch-jobs"), b, nil)
	if err != nil {
		return nil, err
	}

	var res BatchTasksResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
