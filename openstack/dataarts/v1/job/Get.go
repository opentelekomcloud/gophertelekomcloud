package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get is used to view job details.
// Send request GET /v1/{project_id}/jobs/{name}
func Get(client *golangsdk.ServiceClient, jobName, workspace string) (*Job, error) {

	var opts *golangsdk.RequestOpts
	if workspace != "" {
		opts.MoreHeaders = map[string]string{HeaderWorkspace: workspace}
	}

	raw, err := client.Get(client.ServiceURL(jobsEndpoint, jobName), nil, opts)
	if err != nil {
		return nil, err
	}

	var res *Job
	err = extract.Into(raw.Body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
