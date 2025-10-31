package job

import (
	"fmt"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func WaitForJobCompletion(client *golangsdk.ServiceClient, timeout int, jobId string) error {
	return golangsdk.WaitFor(timeout, func() (bool, error) {
		jobStatus, err := ListJobs(client, ListJobsOpts{
			Id: jobId,
		})
		if err != nil {
			return false, err
		}

		time.Sleep(15 * time.Second)

		job := jobStatus.Jobs[0]

		if job.Status == "Completed" {
			return true, nil
		}
		if job.Status == "Failed" {
			return false, fmt.Errorf("job failed: %s", job.FailReason)
		}

		return false, nil
	})
}
