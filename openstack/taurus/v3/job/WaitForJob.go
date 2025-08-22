package job

import (
	"fmt"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func WaitForJobSuccess(client *golangsdk.ServiceClient, secs int, jobID string) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		jobStatus, err := GetJobStatus(client, jobID)
		if err != nil {
			return false, err
		}

		time.Sleep(15 * time.Second)

		job := jobStatus.Job

		if job.Status == "Completed" {
			return true, nil
		}
		if job.Status == "Failed" {
			return false, fmt.Errorf("job failed: %s", job.FailReason)
		}

		return false, nil
	})
}
