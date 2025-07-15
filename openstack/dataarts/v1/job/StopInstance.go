package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// StopInstance is used to stop a specific job instance. A job instance can be stopped only when it is in the running state.
// Send request POST /v1/{project_id}/jobs/{job_name}/instances/{instance_id}/stop
func StopInstance(client *golangsdk.ServiceClient, jobName, instanceId, workspace string) error {
	opts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{},
		OkCodes:     []int{204},
	}
	if workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = workspace
	}

	_, err := client.Post(client.ServiceURL(jobsEndpoint, jobName, instancesEndpoint, instanceId, stopEndpoint), nil, nil, opts)
	if err != nil {
		return err
	}

	return nil
}
