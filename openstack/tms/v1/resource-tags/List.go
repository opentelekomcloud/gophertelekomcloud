package resource_tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ListOpts is the structure that used to query tags detail for specified resource.
type ListOpts struct {
	// Specifies the resource type.
	ResourceType string `q:"resource_type" required:"true"`
	// Specifies the project ID. This parameter is mandatory for region-level resources.
	ProjectId string `q:"project_id"`
}

// List tags of a specific resource.
func List(client *golangsdk.ServiceClient, resourceId string, opts ListOpts) ([]ResourceTag, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resources", resourceId, "tags").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2.0/resources/{resource_id}/tags
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []ResourceTag
	err = extract.IntoSlicePtr(raw.Body, &res, "tags")
	return res, err
}
