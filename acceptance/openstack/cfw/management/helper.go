package management

import (
	"fmt"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	cfwjob "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v3/job"
)

func WaitForJobCompleted(client *golangsdk.ServiceClient, waitTime int, interval time.Duration, jobID string) error {
	jobClient := *client
	jobClient.ResourceBase = jobClient.Endpoint

	return golangsdk.WaitFor(waitTime, func() (bool, error) {
		job, err := cfwjob.Get(client, jobID)
		if err != nil {
			return false, err
		}

		if job.Status == "Success" {
			return true, nil
		}
		if job.Status == "Failed" {
			err = fmt.Errorf("job %s failed", job.Id)
			return false, err
		}

		time.Sleep(interval * time.Second)
		return false, nil
	})
}
