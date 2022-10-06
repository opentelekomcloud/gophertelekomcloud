package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchReplaceDefinerOpts struct {
	Jobs []ReplaceDefinerInfo `json:"jobs"`
}

type ReplaceDefinerInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Whether to replace the definer with the destination database user.
	ReplaceDefiner bool `json:"replace_definer"`
}

func BatchSetDefiner(client *golangsdk.ServiceClient, opts BatchRetryOpts) (*BatchJobsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-replace-definer
	raw, err := client.Post(client.ServiceURL("jobs", "batch-replace-definer"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res BatchJobsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
