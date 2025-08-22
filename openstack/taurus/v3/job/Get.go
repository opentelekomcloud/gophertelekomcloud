package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type JobResponse struct {
	JobID string `json:"job_id"`
}

type JobStatus struct {
	Job Job `json:"job"`
}

type Job struct {
	Status     string `json:"status"`
	JobID      string `json:"id"`
	FailReason string `json:"fail_reason"`
}

func GetJobStatus(client *golangsdk.ServiceClient, jobID string) (*JobStatus, error) {
	raw, err := client.Get(client.ServiceURL("jobs")+"?id="+jobID, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res JobStatus
	return &res, extract.Into(raw.Body, &res)
}
