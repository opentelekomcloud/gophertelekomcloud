package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Stop is used to stop a job.
// Send request POST /v1/{project_id}/jobs/{job_name}/stop
func Stop(client *golangsdk.ServiceClient, jobName, workspace string) error {
	opts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{},
		OkCodes:     []int{204},
	}
	if workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = workspace
	}

	_, err := client.Post(client.ServiceURL(jobsEndpoint, jobName, stopEndpoint), nil, nil, opts)
	if err != nil {
		return err
	}

	return nil
}
