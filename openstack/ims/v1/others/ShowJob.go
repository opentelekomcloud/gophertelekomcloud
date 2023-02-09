package others

import (
	"fmt"
	"net/http"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowJob(client *golangsdk.ServiceClient, jobId string) (*ShowJobResponse, error) {
	// GET /v1/{project_id}/jobs/{job_id}
	raw, err := client.Get(client.ServiceURL(client.ProjectID, "jobs", jobId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowJobResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

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
	// Specifies the job name.
	CurrentTask string `json:"current_task,omitempty"`
	// Specifies the image name.
	ImageName string `json:"image_name,omitempty"`
	// Specifies the job progress.
	ProcessPercent float64 `json:"process_percent,omitempty"`
	// Specifies job execution results.
	Results []JobEntitiesResult `json:"results,omitempty"`
	// Specifies sub-job execution results.
	SubJobsResult []SubJobResult `json:"sub_jobs_result,omitempty"`
	// Specifies the sub-job IDs.
	SubJobsList []string `json:"sub_jobs_list,omitempty"`
}

type JobEntitiesResult struct {
	// Specifies the image ID.
	ImageId string `json:"image_id,omitempty"`
	// Specifies the project ID.
	ProjectId string `json:"project_id,omitempty"`
	// Specifies the job status.
	Status string `json:"status,omitempty"`
}

type SubJobResult struct {
	// Specifies the sub-job status. The value can be:
	//
	// SUCCESS: The sub-job is successfully executed.
	//
	// FAIL: The sub-job failed to be executed.
	//
	// RUNNING: The sub-job is in progress.
	//
	// INIT: The sub-job is being initialized.
	Status string `json:"status,omitempty"`
	// Specifies a sub-job ID.
	JobId string `json:"job_id,omitempty"`
	// Specifies the sub-job type.
	JobType string `json:"job_type,omitempty"`
	// Specifies the start time of the sub-job. The value is in UTC format.
	BeginTime string `json:"begin_time,omitempty"`
	// Specifies the end time of the sub-job. The value is in UTC format.
	EndTime string `json:"end_time,omitempty"`
	// Specifies the error code.
	ErrorCode string `json:"error_code,omitempty"`
	// Specifies the failure cause.
	FailReason string `json:"fail_reason,omitempty"`
	// Specifies the custom attributes of the sub-job. For details, see Table 5.
	//
	// If a sub-job is properly executed, an image ID is returned.
	//
	// If an exception occurs on the sub-job, an error code and associated information are returned.
	Entities SubJobEntities `json:"entities,omitempty"`
}

type SubJobEntities struct {
	// Specifies the image ID.
	ImageId string `json:"image_id,omitempty"`
	// Specifies the image name.
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

func WaitForJob(c *golangsdk.ServiceClient, id string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := ShowJob(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == "SUCCESS" {
			return true, nil
		}

		if current.Status == "FAIL" {
			return false, fmt.Errorf("job failed: %s", current.FailReason)
		}

		return false, nil
	})
}
