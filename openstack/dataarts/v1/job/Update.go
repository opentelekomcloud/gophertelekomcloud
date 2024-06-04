package job

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// Update is used to edit a job.
// Send request PUT /v1/{project_id}/jobs/{job_name}
func Update(client *golangsdk.ServiceClient, jobName string, job Job) error {

	b, err := build.RequestBody(job, "")
	if err != nil {
		return err
	}

	reqOpts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{HeaderContentType: ApplicationJson},
		OkCodes:     []int{204},
	}

	if job.Workspace != "" {
		reqOpts.MoreHeaders = map[string]string{HeaderWorkspace: job.Workspace}
	}

	_, err = client.Put(client.ServiceURL(jobsEndpoint, jobName), b, nil, reqOpts)
	if err != nil {
		return err
	}

	return nil
}
