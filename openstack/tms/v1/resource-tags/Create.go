package resource_tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// BatchOpts is the structure that used to manage tags.
type BatchOpts struct {
	// Specifies the resource list.
	Resources []Resource `json:"resources" required:"true"`
	// Tags list.
	Tags []ResourceTag `json:"tags" required:"true"`
	// Specifies the project ID. This parameter is mandatory when resource_type is a region-specific service.
	ProjectId string `json:"project_id,omitempty"`
}

// Resource is the object that represents the managed resource configuration.
type Resource struct {
	// Specifies the resource type.
	ResourceType string `json:"resource_type" required:"true"`
	// Specifies the resource ID.
	ResourceId string `json:"resource_id" required:"true"`
}

// ResourceTag is the object that represents the tags configuration for batch management.
type ResourceTag struct {
	// Specifies the tag key.
	// The value can contain up to 36 characters including letters, digits, hyphens (-), and underscores (_).
	Key string `json:"key" required:"true"`
	// Specifies the tag value.
	// The value can contain up to 43 characters including letters, digits, periods (.), hyphens (-) and
	// underscores (_). It can be an empty string.
	Value *string `json:"value,omitempty"`
}

// Create is a method to create tags in batch using given parameters.
func Create(client *golangsdk.ServiceClient, opts BatchOpts) ([]FailedResource, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1.0/resource-tags/batch-create
	raw, err := client.Post(client.ServiceURL("resource-tags", "batch-create"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})

	var res []FailedResource
	err = extract.IntoSlicePtr(raw.Body, &res, "failed_resources")
	return res, err
}

// FailedResource is the structure that represents the resource list that set tags failed.
type FailedResource struct {
	// Specifies the resource ID.
	ResourceId string `json:"resource_id"`
	// Specifies the resource type.
	ResourceType string `json:"resource_type"`
	// Specifies the error code.
	ErrorCode string `json:"error_code"`
	// Specifies the error message.
	ErrorMsg string `json:"error_msg"`
}
