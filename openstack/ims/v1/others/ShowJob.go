package others

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Job ID

// GET /v1/{project_id}/jobs/{job_id}

type ShowJobResponse struct {
	// Specifies the job status. The value can be:
	//
	// SUCCESS: The job is successfully executed.
	//
	// FAIL: The job failed to be executed.
	//
	// RUNNING: The job is in progress.
	//
	// INIT: The job is being initialized.
	Status string `json:"status,omitempty"`
	// Specifies the task ID.
	JobId string `json:"job_id,omitempty"`
	// Specifies the job type.
	JobType string `json:"job_type,omitempty"`
	// Specifies the start time of the job. The value is in UTC format.
	BeginTime string `json:"begin_time,omitempty"`
	// Specifies the end time of the job. The value is in UTC format.
	EndTime string `json:"end_time,omitempty"`
	// Specifies the error code.
	ErrorCode string `json:"error_code,omitempty"`
	// Specifies the cause of the failure.
	FailReason string `json:"fail_reason,omitempty"`
	// Specifies the custom attributes of the job.
	//
	// If the job status is normal, the image ID will be returned. If the status is abnormal, an error code and details will be returned.
	Entities JobEntities `json:"entities,omitempty"`
}

type JobEntities struct {
	// Specifies the image ID.
	ImageId string `json:"image_id,omitempty"`

	CurrentTask    string              `json:"current_task,omitempty"`
	ImageName      string              `json:"image_name,omitempty"`
	ProcessPercent float64             `json:"process_percent,omitempty"`
	Results        []JobEntitiesResult `json:"results,omitempty"`
	SubJobsResult  []SubJobResult      `json:"sub_jobs_result,omitempty"`
	SubJobsList    []string            `json:"sub_jobs_list,omitempty"`
}

type JobEntitiesResult struct {
	ImageId   string `json:"image_id,omitempty"`
	ProjectId string `json:"project_id,omitempty"`
	Status    string `json:"status,omitempty"`
}

type SubJobResult struct {
	Status     string         `json:"status,omitempty"`
	JobId      string         `json:"job_id,omitempty"`
	JobType    string         `json:"job_type,omitempty"`
	BeginTime  string         `json:"begin_time,omitempty"`
	EndTime    string         `json:"end_time,omitempty"`
	ErrorCode  string         `json:"error_code,omitempty"`
	FailReason string         `json:"fail_reason,omitempty"`
	Entities   SubJobEntities `json:"entities,omitempty"`
}

type SubJobEntities struct {
	ImageId   string `json:"image_id,omitempty"`
	ImageName string `json:"image_name,omitempty"`
}

func ExtractJobId(err error, raw *http.Response) (*string, error) {
	if err != nil {
		return nil, err
	}

	var res struct {
		// Specifies the asynchronous job ID.
		JobId string `json:"job_id"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.JobId, err
}
