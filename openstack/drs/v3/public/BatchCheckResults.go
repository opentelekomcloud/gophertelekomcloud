package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchQueryPrecheckResultReq struct {
	// Request for querying pre-check results in batches. The value cannot be empty.
	// The values must comply with the UUID rule. The task ID must be unique.
	Jobs []string `json:"jobs"`
}

// POST /v3/{project_id}/jobs/batch-precheck-result

type BatchCheckResultsResponse struct {
	Results []QueryPreCheckResp `json:"results,omitempty"`
	Count   int                 `json:"count,omitempty"`
}

type QueryPreCheckResp struct {
	// Pre-check ID.
	PrecheckId string `json:"precheck_id,omitempty"`
	// Whether the pre-check items are passed. true: indicates that the pre-check is passed.
	// The task can be started only after the pre-check is passed.
	Result bool `json:"result,omitempty"`
	// Pre-check progress, in percentage.
	Process string `json:"process,omitempty"`
	// Percentage of passed pre-checks.
	TotalPassedRate string `json:"total_passed_rate,omitempty"`
	// RDS DB instance ID.
	RdsInstanceId string `json:"rds_instance_id,omitempty"`
	// Migration direction. Values:
	// up: The current cloud is the standby cloud in the cloud DR and backup scenario.
	// down: The current cloud is the active cloud in the cloud DR scenario.
	// non-dbs: self-built databases.
	JobDirection string `json:"job_direction,omitempty"`
	// Pre-check results.
	PrecheckResult []PrecheckResult `json:"precheck_result,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMsg string `json:"error_msg,omitempty"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
}

func BatchCheckResults(client *golangsdk.ServiceClient, opts BatchQueryPrecheckResultReq) (*BatchCheckResultsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/jobs/batch-precheck-result
	raw, err := client.Post(client.ServiceURL("jobs", "batch-precheck-result"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res BatchCheckResultsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type PrecheckResult struct {
	// Check item.
	Item string `json:"item,omitempty"`
	// Check results. Values:
	// PASSED ALARM FAILED
	Result string `json:"result,omitempty"`
	// Failure cause.
	FailedReason string `json:"failed_reason,omitempty"`
	// Encrypted data.
	Data string `json:"data,omitempty"`
	// Row error message.
	RawErrorMsg string `json:"raw_error_msg,omitempty"`
	// Check item group.
	Group string `json:"group,omitempty"`
	// Information about failed subtasks.
	FailedSubJobs []PrecheckFailSubJobVo `json:"failed_sub_jobs,omitempty"`
}

type PrecheckFailSubJobVo struct {
	// ID of the subtask that fails to pass the pre-check.
	Id string `json:"id,omitempty"`
	// The name of the subtask that fails to pass the pre-check.
	Name string `json:"name,omitempty"`
	// Check results.
	CheckResult string `json:"check_result,omitempty"`
}
