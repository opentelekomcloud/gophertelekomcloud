package instances

import (
	"fmt"
	"strings"
	"time"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type JobId struct {
	JobId string `json:"job_id"`
}

// jobNotFoundCode is returned by the job service when a freshly submitted job
// has not been registered yet
const jobNotFoundCode = "DBS.200543"

func WaitForJobCompleted(client *golangsdk.ServiceClient, secs int, jobID string) error {
	jobClient := *client
	jobClient.ResourceBase = jobClient.Endpoint

	return golangsdk.WaitFor(secs, func() (bool, error) {
		job := new(golangsdk.RDSJobStatus)

		requestOpts := &golangsdk.RequestOpts{MoreHeaders: map[string]string{"Content-Type": "application/json"}}
		_, err := jobClient.Get(fmt.Sprintf("%sjobs?id=%s", jobClient.ResourceBase, jobID), job, requestOpts)
		if err != nil {
			if strings.Contains(err.Error(), jobNotFoundCode) {
				time.Sleep(10 * time.Second)
				return false, nil
			}
			return false, err
		}

		if job.Job.Status == "Completed" {
			return true, nil
		}
		if job.Job.Status == "Failed" {
			err = fmt.Errorf("Job failed \n%#v.\n", job.Job)
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
