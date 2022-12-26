package instances

import (
	"fmt"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type JobId struct {
	JobId string `json:"job_id"`
}

func WaitForJobCompleted(client *golangsdk.ServiceClient, secs int, jobID string) error {
	jobClient := *client
	jobClient.ResourceBase = jobClient.Endpoint

	return golangsdk.WaitFor(secs, func() (bool, error) {
		job := new(golangsdk.RDSJobStatus)

		requestOpts := &golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}}
		_, err := jobClient.Get(fmt.Sprintf("%sjobs?id=%s", jobClient.ResourceBase, jobID), job, requestOpts)
		if err != nil {
			return false, err
		}

		if job.Job.Status == "Completed" {
			return true, nil
		}
		if job.Job.Status == "Failed" {
			err = fmt.Errorf("Job failed %s.\n", job.Job.Status)
			return false, err
		}

		time.Sleep(10 * time.Second)
		return false, nil
	})
}

func WaitForStateAvailable(client *golangsdk.ServiceClient, secs int, instanceID string) error {
	jobClient := *client
	jobClient.ResourceBase = jobClient.Endpoint

	return golangsdk.WaitFor(secs, func() (bool, error) {
		job := new(golangsdk.JsonRDSInstanceStatus)

		requestOpts := &golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}}
		_, err := jobClient.Get(fmt.Sprintf("%sinstances?id=%s", jobClient.ResourceBase, instanceID), job, requestOpts)
		if err != nil {
			return false, err
		}

		if job.Instances[0].Status == "ACTIVE" {
			return true, nil
		}
		if job.Instances[0].Status == "FAILED" {
			err = fmt.Errorf("Job failed %s.\n", job.Instances[0].Status)
			return false, err
		}

		return false, nil
	})
}
