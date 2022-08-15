package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func UpdateDDos(client *golangsdk.ServiceClient, floatingIpId string, opts ConfigOpts) (*UpdateResponse, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/antiddos/{floating_ip_id}
	raw, err := client.Put(client.ServiceURL("antiddos", floatingIpId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	var res UpdateResponse
	err = extract.Into(raw, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type UpdateResponse struct {
	// Internal error code
	ErrorCode string `json:"error_code,"`
	// Internal error description
	ErrorDescription string `json:"error_description,"`
	// ID of a task. This ID can be used to query the status of the task.
	// This field is reserved for use in task auditing later. It is temporarily unused.
	TaskId string `json:"task_id,"`
}
