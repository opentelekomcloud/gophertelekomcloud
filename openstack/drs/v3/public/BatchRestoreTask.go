package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchRetryOpts struct {
	Jobs []RetryInfo `json:"jobs"`
}

type RetryInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// This parameter is mandatory and must be set to true.
	IsSyncReEdit bool `json:"is_sync_re_edit,omitempty"`
}

func BatchRestoreTask(client *golangsdk.ServiceClient, opts BatchRetryOpts) (*BatchJobsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-retry-task
	raw, err := client.Post(client.ServiceURL("jobs", "batch-retry-task"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res BatchJobsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
