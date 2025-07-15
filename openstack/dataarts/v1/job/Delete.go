package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteReq struct {
	// Workspace ID.
	Workspace string        `json:"-"`
	Approvers []JobApprover `json:"approvers,omitempty"`
}

// Delete is used to stop a job.
// Send request POST /v1/{project_id}/jobs/{job_name}
func Delete(client *golangsdk.ServiceClient, jobName string, reqOpts *DeleteReq) error {
	b, err := build.RequestBody(reqOpts, "")
	if err != nil {
		return err
	}

	opts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{},
		OkCodes:     []int{204},
	}

	if reqOpts != nil && reqOpts.Workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = reqOpts.Workspace
	}

	_, err = client.DeleteWithBody(client.ServiceURL(jobsEndpoint, jobName), b, opts)
	if err != nil {
		return err
	}

	return nil
}
