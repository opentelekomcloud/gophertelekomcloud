package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const executeImmediatelyEndpoint = "run-immediate"

type ExecuteImmediatelyReq struct {
	// Workspace ID.
	Workspace string `json:"-"`
	// Parameters for executing a job immediately
	JobParams []*JobParam `json:"jobParams,omitempty"`
}

type JobParam struct {
	// Name of the parameter. It cannot exceed 64 characters.
	Name string `json:"name" required:"true"`
	// Value of the parameter. It cannot exceed 1,024 characters.
	Value string `json:"value" required:"true"`
	// Parameter type
	//    variable
	//    constants
	ParamType string `json:"paramType,omitempty"`
}

// ExecuteImmediately is used to execute a job immediately and check whether the job can be executed successfully.
// Send request POST /v1/{project_id}/jobs/{job_name}/run-immediate
func ExecuteImmediately(client *golangsdk.ServiceClient, jobName string, reqOpts *ExecuteImmediatelyReq) (*ExecuteImmediatelyResp, error) {
	b, err := build.RequestBody(reqOpts, "")
	if err != nil {
		return nil, err
	}

	opts := &golangsdk.RequestOpts{}
	if reqOpts.Workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = reqOpts.Workspace
	}
	raw, err := client.Post(client.ServiceURL(jobsEndpoint, jobName, executeImmediatelyEndpoint), b, nil, opts)
	if err != nil {
		return nil, err
	}

	var res ExecuteImmediatelyResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ExecuteImmediatelyResp struct {
	// Job instance ID.
	InstanceId int64 `json:"instanceId"`
}
