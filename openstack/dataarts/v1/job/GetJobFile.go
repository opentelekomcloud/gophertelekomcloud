package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const checkFileEndpoint = "check-file"

type GetJobFileParams struct {
	// If OBS is deployed, the job definition file is stored on OBS, for example, obs://myBucket/jobs.zip.
	Path string `json:"path,omitempty"`
}

// GetJobFile is used to check whether there are jobs and scripts in the job file to be imported from OBS to DLF.
// Send request POST /v1/{project_id}/jobs/check-file
func GetJobFile(client *golangsdk.ServiceClient, opts *GetJobFileParams) (*GetJobFileResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	reqOpts := &golangsdk.RequestOpts{
		OkCodes:     []int{204},
		MoreHeaders: map[string]string{HeaderContentType: ApplicationJson},
	}

	raw, err := client.Post(client.ServiceURL(jobsEndpoint, checkFileEndpoint), b, nil, reqOpts)
	if err != nil {
		return nil, err
	}

	var resp GetJobFileResp
	err = extract.Into(raw.Body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

type GetJobFileResp struct {
	// Job information.
	Jobs []*JobFile `json:"jobs,omitempty"`
	// Script information.
	Scripts []*Script `json:"scripts,omitempty"`
}

type JobFile struct {
	// Job parameter.
	Params map[string]string `json:"params,omitempty"`
	// Job name.
	Name string `json:"name"`
	// Path of the job.
	Path string `json:"path"`
}

type Script struct {
	// Script name.
	Name string `json:"name"`
	// Path of the script.
	Path string `json:"path"`
}
