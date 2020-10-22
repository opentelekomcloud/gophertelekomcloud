package instances

import (
	"fmt"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type DeleteInstanceRdsResponse struct {
	JobId string `json:"job_id"`
}

type EnlargeVolumeResp struct {
	JobId string `json:"job_id"`
}

type RestartRdsResponse struct {
	JobId string `json:"job_id"`
}

type SingleToHaResponse struct {
	JobId string `json:"job_id"`
}

type ResizeFlavor struct {
	JobId string `json:"job_id"`
}

type CreateRds struct {
	Instance Instance `json:"instance"`
	JobId    string   `json:"job_id"`
	OrderId  string   `json:"order_id"`
}

func WaitForJobCompleted(client *golangsdk.ServiceClient, secs int, jobID string) error {
	jobClient := *client
	jobClient.ResourceBase = jobClient.Endpoint

	return golangsdk.WaitFor(secs, func() (bool, error) {
		job := new(golangsdk.JobStatus)
		_, err := jobClient.Get(jobClient.ServiceURL("jobs", jobID), &job, nil)
		time.Sleep(5 * time.Second)
		if err != nil {
			return false, err
		}

		if job.Status == "Completed" {
			return true, nil
		}
		if job.Status == "Failed" {
			err = fmt.Errorf("Job failed with code %s: %s.\n", job.ErrorCode, job.FailReason)
			return false, err
		}

		return false, nil
	})
}
