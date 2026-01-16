package resourcegroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// DeleteOpts contains the options for batch deleting resource groups.
type DeleteOpts struct {
	// Specifies the list of resource group IDs to delete.
	// A maximum of 100 resource groups can be deleted at a time.
	GroupIds []string `json:"group_ids" required:"true"`
}

// DeleteResponse contains the response from the batch delete request.
type DeleteResponse struct {
	// Specifies the list of deleted resource group IDs.
	GroupIds []string `json:"group_ids"`
}

// Delete batch deletes resource groups.
func Delete(client *golangsdk.ServiceClient, opts DeleteOpts) (*DeleteResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/resource-groups/batch-delete
	raw, err := client.Post(client.ServiceURL("resource-groups", "batch-delete"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res DeleteResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
