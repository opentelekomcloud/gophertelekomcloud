package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const restartEndpoint = "restart"

// RetryInstance is used to retry a specific job instance. A job instance can be retried only when it is in the successful, failed, or canceled state.
// Send request POST /v1/{project_id}/jobs/{job_name}/instances/{instance_id}/restart
func RetryInstance(client *golangsdk.ServiceClient, jobName, instanceId, workspace string) error {
	opts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{},
		OkCodes:     []int{204},
	}
	if workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = workspace
	}

	_, err := client.Post(client.ServiceURL(jobsEndpoint, jobName, instancesEndpoint, instanceId, restartEndpoint), nil, nil, opts)
	if err != nil {
		return err
	}

	return nil
}
