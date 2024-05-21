package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetJobInstance is used to view job instance details, including the execution information about each node in a job instance.
// Send request GET /v1/{project_id}/jobs/{job_name}/instances/{instance_id}
func GetJobInstance(client *golangsdk.ServiceClient, jobName, instanceId, workspace string) (*JobInstanceResp, error) {

	var opts *golangsdk.RequestOpts
	if workspace != "" {
		opts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	raw, err := client.Get(client.ServiceURL(jobsEndpoint, jobName, instancesEndpoint, instanceId), nil, opts)
	if err != nil {
		return nil, err
	}

	var res JobInstanceResp
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type JobInstanceResp struct {
	// Job name.
	JobName string `json:"jobName"`
	// Job instance ID.
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
	// Total number of node records.
	Total int `json:"total"`
	// Node instance status.
	Nodes []NodeJobInstanceResp `json:"nodes"`
	// Whether the job instance status is forcibly successful
	//
	// Default value: false
	ForceSuccess bool `json:"forceSuccess,omitempty"`
	// Whether the job instance status is failure ignored
	//
	// Default value: false
	IgnoreSuccess bool `json:"ignoreSuccess,omitempty"`
}

type NodeJobInstanceResp struct {
	// Node name.
	NodeName string `json:"nodeName"`
	// Node status.
	//    waiting
	//    running
	//    success
	//    fail
	//    skip
	//    pause
	//    manual-stop
	Status string `json:"status,omitempty"`
	// Planned execution time of the job instance.
	PlanTime int64 `json:"planTime"`
	// Actual execution start time of the job instance.
	StartTime int64 `json:"startTime"`
	// Actual execution end time of the job instance.
	EndTime int64 `json:"endTime,omitempty"`
	// Node type.
	Type string `json:"type"`
	// Number of attempts upon a failure.
	RetryTimes int `json:"retryTimes,omitempty"`
	// Job instance ID.
	InstanceId int64 `json:"instanceId"`
	// Rows of input data.
	InputRowCount int64 `json:"inputRowCount,omitempty"`
	// Write speed (row/second)
	Speed float64 `json:"speed,omitempty"`
	// Path for storing node execution logs.
	LogPath string `json:"logPath,omitempty"`
}
