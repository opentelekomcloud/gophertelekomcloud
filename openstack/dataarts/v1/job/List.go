package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Offset  int    `q:"offset"`
	Limit   int    `q:"limit"`
	JobType string `q:"jobType"`
	JobName string `q:"jobName"`
}

// List is used to query a list of batch or real-time jobs. A maximum of 100 jobs can be returned for each query.
// Send request GET /v1/{project_id}/jobs?jobType={jobType}&offset={offset}&limit={limit}&jobName={jobName}
func List(client *golangsdk.ServiceClient, opts ListOpts, workspace string) (*ListResp, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints(jobsEndpoint).WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	var reqOpts golangsdk.RequestOpts
	if workspace != "" {
		reqOpts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, &reqOpts)
	if err != nil {
		return nil, err
	}

	var res ListResp
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type ListResp struct {
	// Total is the number of jobs.
	Total int `json:"total"`
	// Jobs is a list of jobs.
	Jobs []*JobList `json:"jobs"`
}

type JobList struct {
	// Job name.
	Name string `json:"name"`
	// Job type.
	//    REAL_TIME: real-time processing
	//    BATCH: batch processing
	JobType string `json:"jobType"`
	// Job owner. The length cannot exceed 128 characters.
	Owner string `json:"owner,omitempty"`
	// Job priority. The value ranges from 0 to 2. The default value is 0. 0 indicates a top priority, 1 indicates a medium priority, and 2 indicates a low priority.
	Priority int `json:"priority,omitempty"`
	// Job status.
	// When jobType is set to REAL_TIME, the status is as follows:
	//    STARTING
	//    NORMAL
	//    EXCEPTION
	//    STOPPING
	//    STOPPED
	//
	// When jobType is set to BATCH, the status is as follows:
	//    SCHEDULING
	//    STOPPED
	//    PAUSED
	Status string `json:"status"`
	// Job creator.
	CreateUser string `json:"createUser"`
	// Time when the job is created.
	CreateTime int64 `json:"createTime"`
	// Time when the job starts running.
	StartTime int64 `json:"startTime,omitempty"`
	// Time when the job stops running.
	EndTime int64 `json:"endTime,omitempty"`
	// Most recent running status of the job instance. This parameter is available only when jobType is set to BATCH.
	LastInstanceStatus string `json:"lastInstanceStatus,omitempty"`
	// Time when the most recent job instance stops to run. This parameter is available only when jobType is set to BATCH.
	LastInstanceEndTime int64 `json:"lastInstanceEndTime,omitempty"`
	// Time when the job was last updated.
	LastUpdateTime int64 `json:"lastUpdateTime,omitempty"`
	// User who last updated the job.
	LastUpdateUser string `json:"lastUpdateUser,omitempty"`
	// Path of the job.
	Path string `json:"path,omitempty"`
	// Whether the job is a single-task job.
	SingleNodeJobFlag bool `json:"singleNodeJobFlag,omitempty"`
}
