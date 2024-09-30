package protectedinstances

import (
	"fmt"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type JobStatus struct {
	Status     string    `json:"status"`
	Entities   JobEntity `json:"entities"`
	JobID      string    `json:"job_id"`
	JobType    string    `json:"job_type"`
	BeginTime  string    `json:"begin_time"`
	EndTime    string    `json:"end_time"`
	ErrorCode  string    `json:"error_code"`
	FailReason string    `json:"fail_reason"`
	Message    string    `json:"message"`
	Code       string    `json:"code"`
}

type JobEntity struct {
	InstanceID string   `json:"protected_instance_id"`
	SubJobs    []SubJob `json:"sub_jobs"`
}

type SubJob struct {
	// Specifies the task ID.
	ID string `json:"job_id"`
	// Task type.
	Type string `json:"job_type"`
	// Specifies the task status.
	Status string `json:"status"`
	// Specifies the time when the task started.
	BeginTime string `json:"begin_time"`
	// Specifies the time when the task finished.
	EndTime string `json:"end_time"`
	// Specifies the returned error code when the task execution fails.
	ErrorCode string `json:"error_code"`
	// Specifies the cause of the task execution failure.
	FailReason string `json:"fail_reason"`
	// Specifies the object of the task.
	Entities map[string]string `json:"entities"`
}

func WaitForJobSuccess(client *golangsdk.ServiceClient, secs int, jobID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		job := new(JobStatus)
		_, err := client.Get(client.ServiceURL("jobs", jobID), &job, nil)
		if err != nil {
			return false, err
		}
		time.Sleep(5 * time.Second)

		if job.Status == "SUCCESS" {
			return true, nil
		}
		if job.Status == "FAIL" {
			err = fmt.Errorf("job failed with code %s: %s", job.ErrorCode, job.FailReason)
			return false, err
		}

		return false, nil
	})
}

func GetJobEntity(client *golangsdk.ServiceClient, jobID string, label string) (interface{}, error) {
	if label != "protected_instance_id" {
		return nil, fmt.Errorf("unsupported label %s in GetJobEntity", label)
	}

	job := new(JobStatus)
	_, err := client.Get(client.ServiceURL("jobs", jobID), &job, nil)
	if err != nil {
		return nil, err
	}

	if job.Status == "SUCCESS" {
		if instanceID := job.Entities.InstanceID; instanceID != "" {
			return instanceID, nil
		}
	}

	return nil, fmt.Errorf("unexpected conversion error in GetJobEntity")
}
