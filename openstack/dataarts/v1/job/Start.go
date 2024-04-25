package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

const startEndpoint = "start"

type StartReq struct {
	// Parameter for starting the job.
	JobParams []*JobParam `json:"jobParams,omitempty"`
}

// Start is used to start a job.
// Send request POST /v1/{project_id}/jobs/{job_name}/start
func Start(client *golangsdk.ServiceClient, jobName, workspace string, reqOpts *StartReq) error {
	b, err := build.RequestBody(reqOpts, "")
	if err != nil {
		return err
	}

	opts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			HeaderContentType: ApplicationJson,
		},
		OkCodes: []int{204},
	}
	if workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = workspace
	}

	_, err = client.Post(client.ServiceURL(jobsEndpoint, jobName, startEndpoint), b, nil, opts)
	if err != nil {
		return err
	}

	return nil
}
