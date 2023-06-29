package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchStartJobOpts struct {
	// Request list for starting tasks in batches.
	Jobs []StartInfo `json:"jobs"`
}

type StartInfo struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Task start time. The timestamp is accurate to milliseconds, for example, 1608188903063.
	// If the value is empty, the task is started immediately.
	StartTime string `json:"start_time,omitempty"`
}

func BatchStartTasks(client *golangsdk.ServiceClient, opts BatchStartJobOpts) (*BatchTasksResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-starting
	raw, err := client.Post(client.ServiceURL("jobs", "batch-starting"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202}})
	if err != nil {
		return nil, err
	}

	var res BatchTasksResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BatchTasksResponse struct {
	Results []IdJobResp `json:"results,omitempty"`
	Count   int         `json:"count,omitempty"`
}

type IdJobResp struct {
	// Task ID.
	Id     string `json:"id"`
	Status string `json:"status,omitempty"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMsg string `json:"error_msg,omitempty"`
}
