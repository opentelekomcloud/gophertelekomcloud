package job

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateSpecificOpts struct {
	// Jobs is a jobs list.
	Jobs []Job `json:"jobs" required:"true"`
}

// CreateSpecific is used to create a job in a specified cluster.
// Send request /v1.1/{project_id}/clusters/{cluster_id}/cdm/job
func CreateSpecific(client *golangsdk.ServiceClient, clusterID string, opts CreateSpecificOpts) (*SpecificResp, error) {

	reqOpts := &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			HeaderContentType: ApplicationJson,
		},
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL(clustersEndpoint, clusterID, cdmEndpoint, jobEndpoint), b, nil, reqOpts)

	if err != nil {
		return nil, err
	}

	var resp SpecificResp
	err = extract.Into(raw.Body, &resp)
	return &resp, err
}

// SpecificResp holds job running information
type SpecificResp struct {
	// Name is a job name.
	Name string `json:"name"`
	// ValidationResult is an array of ValidationResult objects.
	ValidationResult []*JobValidationResult `json:"validation-result"`
}

type JobValidationResult struct {
	// Message is an error message.
	Message string `json:"message,omitempty"`
	// Status can be ERROR,WARNING.
	Status string `json:"status,omitempty"`
}
