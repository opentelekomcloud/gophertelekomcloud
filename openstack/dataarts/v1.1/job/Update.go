package job

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Jobs []Job `json:"jobs" required:"true"`
}

// Update is used to modify a job.
// Send request PUT /v1/{project_id}/connections/{connection_name}?ischeck=true
func Update(client *golangsdk.ServiceClient, clusterId, jobName string, opts *UpdateOpts) (*UpdateResp, error) {

	b, err := build.RequestBody(opts, "jobs")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(clustersEndpoint, clusterId, cdmEndpoint, jobEndpoint, jobName), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp UpdateResp
	err = extract.Into(raw.Body, &resp)
	return &resp, err

}

type UpdateResp struct {
	// Submissions is an array of StartJobSubmission objects.
	ValidationResult []JobValidationResult `json:"validation-result"`
}
