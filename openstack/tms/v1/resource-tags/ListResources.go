package resource_tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListResourceOpts struct {
	// Resource type. This parameter is case-sensitive.
	// Supported resource types can be provided as ecs,scaling_group,
	// images, disk, vpcs, security-groups, shared_bandwidth, eip, cdn.
	ResourceTypes []string `json:"resource_types" required:"true"`
	// Specifies the project ID. This parameter is mandatory for region-level resources.
	ProjectId string `json:"project_id"`
	// Tags
	Tags []ListResourceTag `json:"tags" required:"true"`
	// Specifies whether to query only untagged resources. If this parameter is set to true,
	// only untagged resources are queried.
	WithoutAnyTag *bool `json:"without_any_tag,omitempty"`
	// Index position. The query starts from the next data specified by offset.
	// The value must be a number and cannot be negative. The default value is 0.
	Offset *int `json:"offset,omitempty"`
	// The maximum queries supported. The value 200 is used by default if this parameter is not set.
	// The value range is 1 to 200.
	Limit *int `json:"limit,omitempty"`
}

type ListResourceTag struct {
	// Specifies the tag key.
	// The value can contain up to 36 characters including letters, digits, hyphens (-), and underscores (_).
	Key string `json:"key" required:"true"`
	// Specifies tag values.
	Values []string `json:"values" required:"true"`
}

// ListResources filter resources by tag.
func ListResources(client *golangsdk.ServiceClient, opts ListResourceOpts) ([]ResourceResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1.0/resource-instances/filter
	raw, err := client.Post(client.ServiceURL("resource-instances", "filter"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json", "X-Language": "en-us"},
	})

	var res []ResourceResp
	err = extract.IntoSlicePtr(raw.Body, &res, "resources")
	return res, err
}

type ResourceResp struct {
	// ProjectID
	ProjectId string `json:"project_id"`
	// Specifies the project name.
	ProjectName string `json:"project_name"`
	// Specifies the resource ID.
	ResourceId string `json:"resource_id"`
	// Specifies the resource name.
	ResourceName string `json:"resource_name"`
	// Specifies the resource type.
	ResourceType string `json:"resource_type"`
	// Specifies the resource tag.
	Tags []ResourceTag `json:"tags"`
}
