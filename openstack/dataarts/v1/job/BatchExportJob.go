package job

import (
	"io"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

const batchExportEndpoint = "batch-export"

type BatchExportReq struct {
	// Workspace ID.
	Workspace string `json:"-"`
	// A list of jobs to be exported. A maximum of 100 jobs can be exported at a time.
	JobList []string `json:"jobList" required:"true"`
	// Specifies whether to export the scripts and resources that the job depends on.
	// Default value: true
	ExportDepend bool `json:"exportDepend,omitempty"`
}

// BatchExportJob is used to batch export jobs, including job dependency scripts and CDM job definitions.
// Send request POST /v1/{project_id}/jobs/batch-export
func BatchExportJob(client *golangsdk.ServiceClient, reqOpts *BatchExportReq) (io.ReadCloser, error) {
	b, err := build.RequestBody(reqOpts, "")
	if err != nil {
		return nil, err
	}

	opts := &golangsdk.RequestOpts{}
	if reqOpts.Workspace != "" {
		opts.MoreHeaders[HeaderWorkspace] = reqOpts.Workspace
	}

	raw, err := client.Post(client.ServiceURL(jobsEndpoint, batchExportEndpoint), b, nil, opts)
	if err != nil {
		return nil, err
	}

	return raw.Body, err
}
