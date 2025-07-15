package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const importEndpoint = "import"

type ImportReq struct {
	// Workspace ID.
	Workspace string `json:"-"`
	// If OBS is deployed, it refers to the OBS path for storing the job definition file. For the format of the job definition file, see the response message of the exported job, for example, obs://myBucket/jobs.zip.
	Path string `json:"path" required:"true"`
	// Public job parameter.
	Params map[string]string `json:"params,omitempty"`
	// Policy for specifying how to handle duplicate names. The options are as follows:
	//    SKIP
	//    OVERWRITE
	// Default value: SKIP
	SameNamePolicy string `json:"sameNamePolicy,omitempty"`
	// Job parameter.
	JobsParam []*JobParamImported `json:"jobsParam,omitempty"`
	// User that executes the job.
	ExecuteUser string `json:"executeUser,omitempty"`
	// This parameter is required if the review function is enabled. It indicates the target status of the job. The value can be SAVED, SUBMITTED, or PRODUCTION.
	//    SAVED indicates that the job is saved but cannot be scheduled or executed. It can be executed only after submitted and approved.
	//    SUBMITTED indicates that the job is automatically submitted after it is saved and can be executed after it is approved.
	//    PRODUCTION indicates that the job can be directly executed after it is created. Note: Only the workspace administrator can create jobs in the PRODUCTION state.
	TargetStatus string `json:"targetStatus,omitempty"`
	// Job approver. This parameter is required if the review function is enabled.
	Approvers []*JobApprover `json:"approvers,omitempty"`
}

type JobParamImported struct {
	Name   string            `json:"name" required:"true"`
	Params map[string]string `json:"params,omitempty"`
}

// ImportJob is used to import one or more job files from OBS to DLF.
// Send request POST /v1/{project_id}/jobs/import
func ImportJob(client *golangsdk.ServiceClient, reqOpts ImportReq) (*ImportResp, error) {
	b, err := build.RequestBody(reqOpts, "")
	if err != nil {
		return nil, err
	}

	opts := &golangsdk.RequestOpts{
		OkCodes: []int{200},
	}
	if reqOpts.Workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = reqOpts.Workspace
	}
	raw, err := client.Post(client.ServiceURL(jobsEndpoint, importEndpoint), b, nil, opts)
	if err != nil {
		return nil, err
	}

	var res ImportResp
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type ImportResp struct {
	// ID of the task. Used to call the API for querying system tasks to obtain the import status.
	TaskId string `json:"taskId" required:"true"`
}
