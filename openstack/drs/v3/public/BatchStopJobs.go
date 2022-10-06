package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchPauseJobOpts struct {
	// The value cannot contain empty objects. The value of job_id must comply with the UUID rule.
	Jobs []PauseInfo `json:"jobs"`
}

type PauseInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Pause type. target: Stop replay. all: Stop log capturing and replay.
	// Values: target all
	PauseMode string `json:"pause_mode"`
}

func BatchStopJobs(client *golangsdk.ServiceClient, opts BatchPauseJobOpts) (*BatchJobsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-pause-task
	raw, err := client.Post(client.ServiceURL("jobs", "batch-pause-task"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res BatchJobsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
