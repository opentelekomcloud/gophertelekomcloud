package link

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	Links []*Link `json:"links" required:"true"`
}

// Update is used to modify a link.
// Send request PUT /v1.1/{project_id}/clusters/{cluster_id}/cdm/link/{link_name}
func Update(client *golangsdk.ServiceClient, clusterId, linkName string, opts *UpdateOpts) (*UpdateResp, error) {

	b, err := build.RequestBody(opts, "links")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(clustersEndpoint, clusterId, cdmEndpoint, linkEndpoint, linkName), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp *UpdateResp
	err = extract.Into(raw.Body, resp)
	return resp, err

}

type UpdateResp struct {
	// Submissions is an array of LinkValidationResult objects.
	ValidationResult []*LinkValidationResult `json:"validation-result"`
}
