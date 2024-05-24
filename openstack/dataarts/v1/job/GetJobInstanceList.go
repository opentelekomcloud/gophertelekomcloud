package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const detailEndpoint = "detail"

type GetJobInstanceListReq struct {
	// Workspace ID.
	Workspace string `json:"-"`
	// Project ID.
	ProjectId string `q:"project_id" required:"true"`
	// Job name.
	//    If you want to query the instance list of a specific batch job, jobName is the batch job name.
	//    If you want to query sub-jobs associated with a node in a real-time job, the jobName format is real-time job name _ node name.
	JobName string `q:"jobName"`
	// Minimum planned job execution time in milliseconds. Job instances whose start time is greater than this time are returned.
	MinPlanTime int64 `q:"minPlanTime"`
	// Minimum planned job execution time in milliseconds. Job instances whose start time is less than this time are returned.
	MaxPlanTime int64 `q:"maxPlanTime"`
	// The maximum number of records on each page.
	// The parameter value ranges from 1 to 1000.
	// Default value: 10
	Limit int `q:"limit"`
	// Start page of the paging list. The default value is 0. The value must be greater than or equal to 0.
	Offset int `q:"offset"`
	// Job instance status.
	//    waiting
	//    running
	//    success
	//    fail
	//    running-exception
	//    pause
	//    manual-stop
	Status string `q:"status"`
	// Job scheduling modes.
	//    0: General scheduling
	//    2: Manual scheduling
	//    5: PatchData
	//    6: Subjob scheduling
	//    7: Schedule once
	InstanceType int `q:"instanceType"`
	// Whether exact search of jobs by name is supported
	PreciseQuery bool `q:"preciseQuery"`
}

// GetJobInstanceList is used to view a job instance list.
//
// A job instance is generated each time you run a batch job for which periodic scheduling or event-based scheduling is configured.
// If a real-time job contains a node for which periodic scheduling or event-based scheduling is configured, you can call this API to view the instance list of the subjobs associated with the node.
// The format of the jobName parameter is real-time job name_node name.
//
// Send request GET  /v1/{project_id}/jobs/instances/detail?jobName={jobName}&minPlanTime={minPlanTime}&maxPlanTime={maxPlanTime}&limit={limit}&offset={offset}&status={status}&instanceType={instanceType}&preciseQuery={preciseQuery}
func GetJobInstanceList(client *golangsdk.ServiceClient, reqOpts GetJobInstanceListReq) (*GetJobInstanceListResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(jobsEndpoint, instancesEndpoint, detailEndpoint).WithQueryParams(&reqOpts).Build()
	if err != nil {
		return nil, err
	}
	var opts golangsdk.RequestOpts
	if reqOpts.Workspace != "" {
		opts.MoreHeaders = map[string]string{HeaderWorkspace: reqOpts.Workspace}
	}

	raw, err := client.Get(url.String(), nil, &opts)

	if err != nil {
		return nil, err
	}

	var res GetJobInstanceListResp
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type GetJobInstanceListResp struct {
	// Total number of records.
	Total int `json:"total"`
	// Job instance status.
	Instances []*NodeStatusListResp `json:"instances"`
}

type NodeStatusListResp struct {
	// Job name. When you view the instance list of a specified batch job, jobName is the name of the batch job.
	// When you view the subjobs associated with a node in a real-time job, jobName is in format of real-time job name_node name.
	JobName string `json:"jobName"`
	// Job ID
	JobId string `json:"jobId,omitempty"`
	// Name of a job instance recorded by the log, rather than the name defined during job creation
	JobInstanceName string `json:"jobInstanceName"`
	// Job instance ID, which is used to query job instance details.
	InstanceId int64 `json:"instanceId"`
	// Job instance status.
	//    waiting
	//    running
	//    success
	//    fail
	//    running-exception
	//    pause
	//    manual-stop
	Status string `json:"status"`
	// Planned execution time of the job instance.
	PlanTime int64 `json:"planTime"`
	// Actual execution start time of the job instance.
	StartTime int64 `json:"startTime"`
	// Actual execution end time of the job instance.
	EndTime int64 `json:"endTime,omitempty"`
	// Execution duration in milliseconds.
	ExecuteTime int64 `json:"executeTime,omitempty"`

	// Time when a job is submitted.
	SubmitTime int64 `json:"submitTime"`
	// Node instance status.
	// Job scheduling modes.
	//    0: General scheduling
	//    2: Manual scheduling
	//    5: PatchData
	//    6: Subjob scheduling
	//    7: Schedule once
	InstanceType int `json:"instanceType"`
	// Whether the job instance status is forcibly successful
	//
	// Default value: false
	ForceSuccess bool `json:"forceSuccess,omitempty"`
	// Whether the job instance status is failure ignored
	//
	// Default value: false
	IgnoreSuccess bool `json:"ignoreSuccess,omitempty"`
	// Job instance version
	Version bool `json:"version,omitempty"`
}
