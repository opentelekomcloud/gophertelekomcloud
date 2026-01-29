package resourcegroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// BatchDeleteResourcesOpts contains the options for batch deleting resources from a resource group.
type BatchDeleteResourcesOpts struct {
	// Specifies the list of resources to delete.
	// A maximum of 1000 resources can be deleted at a time.
	Resources []ResourceItem `json:"resources" required:"true"`
}

// BatchDeleteResourcesResponse contains the response from the BatchDeleteResources request.
type BatchDeleteResourcesResponse struct {
	// Specifies the number of resources that were deleted successfully.
	SucceedCount int `json:"succeed_count"`
}

// BatchDeleteResources removes resources from a resource group in batches.
func BatchDeleteResources(client *golangsdk.ServiceClient, groupId string, opts BatchDeleteResourcesOpts) (*BatchDeleteResourcesResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/resource-groups/{group_id}/resources/batch-delete
	raw, err := client.Post(client.ServiceURL("resource-groups", groupId, "resources", "batch-delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res BatchDeleteResourcesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
