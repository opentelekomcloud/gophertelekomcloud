package job

import (
	"io"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const exportEndpoint = "export"

// ExportJob is used to export a job, including job definitions, job dependency scripts, and CDM job definitions.
// Send request POST /v1/{project_id}/jobs/{job_name}/export
func ExportJob(client *golangsdk.ServiceClient, jobName, workspace string) (io.ReadCloser, error) {

	opts := &golangsdk.RequestOpts{}
	if workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = workspace
	}
	raw, err := client.Post(client.ServiceURL(jobsEndpoint, jobName, exportEndpoint), nil, nil, opts)
	if err != nil {
		return nil, err
	}

	return raw.Body, err
}
