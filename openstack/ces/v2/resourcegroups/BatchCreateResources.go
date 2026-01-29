package resourcegroups

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ResourceDimension represents a resource dimension.
type ResourceDimension struct {
	// Specifies the dimension name.
	// It can contain 1 to 32 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	Name string `json:"name" required:"true"`
	// Specifies the dimension value, which is the resource instance ID.
	// It can contain 1 to 256 characters.
	Value string `json:"value" required:"true"`
}

// ResourceItem represents a resource to be added to or deleted from a resource group.
type ResourceItem struct {
	// Specifies the namespace of a service.
	// The value must be in the service.item format and can contain 3 to 32 characters.
	// service and item must start with a letter and contain only letters, digits, and underscores (_).
	Namespace string `json:"namespace" required:"true"`
	// Specifies the resource dimension information.
	// A maximum of 4 dimensions can be specified.
	Dimensions []ResourceDimension `json:"dimensions" required:"true"`
}

// BatchCreateResourcesOpts contains the options for batch adding resources to a resource group.
type BatchCreateResourcesOpts struct {
	// Specifies the list of resources to add.
	// A maximum of 1000 resources can be added at a time.
	Resources []ResourceItem `json:"resources" required:"true"`
}

// BatchCreateResourcesResponse contains the response from the BatchCreateResources request.
type BatchCreateResourcesResponse struct {
	// Specifies the number of resources that were added successfully.
	SucceedCount int `json:"succeed_count"`
}

// BatchCreateResources adds resources to a resource group in batches.
func BatchCreateResources(client *golangsdk.ServiceClient, groupId string, opts BatchCreateResourcesOpts) (*BatchCreateResourcesResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/resource-groups/{group_id}/resources/batch-create
	raw, err := client.Post(client.ServiceURL("resource-groups", groupId, "resources", "batch-create"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res BatchCreateResourcesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
