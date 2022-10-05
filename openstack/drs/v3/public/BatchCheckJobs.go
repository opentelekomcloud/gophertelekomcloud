package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type PreCheckInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Pre-check mode. Value: forStartJob
	PrecheckMode string `json:"precheck_mode"`
}

func BatchCheckJobs(client *golangsdk.ServiceClient, opts []PreCheckInfo) (*BatchCheckJobsResponse, error) {
	b, err := build.RequestBody(opts, "jobs")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-precheck
	raw, err := client.Post(client.ServiceURL("jobs", "batch-precheck"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res BatchCheckJobsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchCheckJobsResponse struct {
	Results []PostPreCheckResp `json:"results,omitempty"`
	Count   int32              `json:"count,omitempty"`
}

type PostPreCheckResp struct {
	// Task ID
	Id string `json:"id,omitempty"`
	// Pre-check ID.
	PrecheckId string `json:"precheck_id,omitempty"`
	// Success or failure status.
	// Values: success failed
	Status string `json:"status,omitempty"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMsg string `json:"error_msg,omitempty"`
}
