package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchPreCheckReq struct {
	Jobs []PreCheckInfo `json:"jobs" required:"true"`
}

type PreCheckInfo struct {
	JobId        string `json:"job_id" required:"true"`
	PreCheckMode string `json:"precheck_mode" required:"true"`
}

type QueryPreCheckResultReq struct {
	Jobs []string `json:"jobs" required:"true"`
}

func BatchCheckTasks(client *golangsdk.ServiceClient, opts BatchPreCheckReq) (*BatchCheckTasksResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-precheck
	raw, err := client.Post(client.ServiceURL("jobs", "batch-precheck"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res BatchCheckTasksResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchCheckTasksResponse struct {
	Results []PreCheckDetail `json:"results"`
	Count   int              `json:"count"`
}

type PreCheckDetail struct {
	Id         string `json:"id"`
	PreCheckId string `json:"precheck_id"`
	Status     string `json:"status"`
	ErrorCode  string `json:"error_code"`
	ErrorMsg   string `json:"error_msg"`
}
