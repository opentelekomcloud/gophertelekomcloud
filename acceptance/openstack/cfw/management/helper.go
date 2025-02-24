package management

import (
	"fmt"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	cfwjob "github.com/opentelekomcloud/gophertelekomcloud/openstack/cfw/v3/job"
)

func waitForJobCompleted(client *golangsdk.ServiceClient, secs int, jobID string) error {
	jobClient := *client
	jobClient.ResourceBase = jobClient.Endpoint

	return golangsdk.WaitFor(secs, func() (bool, error) {
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

		time.Sleep(5 * time.Second)
		return false, nil
	})
}
