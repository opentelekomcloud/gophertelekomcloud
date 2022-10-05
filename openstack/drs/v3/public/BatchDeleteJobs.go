package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchDeleteJobReq struct {
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
	Count   int32       `json:"count,omitempty"`
}

func BatchDeleteJobs(client *golangsdk.ServiceClient, opts BatchDeleteJobReq) (*BatchJobsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// DELETE /v3/{project_id}/jobs/batch-jobs
	raw, err := client.DeleteWithBody(client.ServiceURL("jobs", "batch-jobs"), b, nil)
	if err != nil {
		return nil, err
	}

	var res BatchJobsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
