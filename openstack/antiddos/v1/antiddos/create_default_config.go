package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ConfigOpts struct {
	// Whether to enable L7 defense
	EnableL7 bool `json:"enable_L7"`
	// Position ID of traffic. The value ranges from 1 to 9.
	TrafficPosId int `json:"traffic_pos_id"`
	// Position ID of number of HTTP requests. The value ranges from 1 to 15.
	HttpRequestPosId int `json:"http_request_pos_id"`
	// Position ID of access limit during cleaning. The value ranges from 1 to 8.
	CleaningAccessPosId int `json:"cleaning_access_pos_id"`
	// Application type ID. Possible values: 0 1
	AppTypeId int `json:"app_type_id"`
}

func CreateDefaultConfig(client *golangsdk.ServiceClient, opts ConfigOpts) (*TaskResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/antiddos/default-config
	raw, err := client.Post(client.ServiceURL("antiddos", "default-config"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res TaskResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type TaskResponse struct {
	// Internal error code
	ErrorCode string `json:"error_code"`
	// Internal error description
	ErrorMessage string `json:"error_msg"`
	// ID of a task. This ID can be used to query the status of the task.
	// This field is reserved for use in task auditing later. It is temporarily unused.
	TaskId string `json:"task_id"`
}
