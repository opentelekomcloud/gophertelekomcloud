package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ShowNewTaskStatusOpts struct {
	// Task ID (non-negative integer) character string
	TaskId string `q:"task_id"`
}

func ShowNewTaskStatus(client *golangsdk.ServiceClient, opts ShowNewTaskStatusOpts) (*ShowNewTaskStatusResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/query_task_status
	raw, err := client.Get(client.ServiceURL("query_task_status")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowNewTaskStatusResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowNewTaskStatusResponse struct {
	// Status of a task, which can be one of the following: success, failed, waiting, running, preprocess, ready
	TaskStatus string `json:"task_status"`
	// Additional information about a task
	TaskMsg string `json:"task_msg"`
}
