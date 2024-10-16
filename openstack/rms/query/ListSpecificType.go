package query

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListSpecificOpts struct {
	// Specifies the number of records returned on each page during pagination query.
	Limit *int `q:"limit"`
	// Specifies the start flag for querying the current page. If this parameter is left blank, the first page is queried.
	// The marker for querying the next page is the next_marker in the page_info object returned on the current page.
	Marker string `q:"marker"`
	// Specifies the region ID.
	RegionId string `q:"region_id"`
	// Specifies the tag.
	Tag map[string]string `q:"tag"`
}

func ListSpecificType(client *golangsdk.ServiceClient, domainId, service, resourceType string, opts ListSpecificOpts) ([]Resource, error) {
	// GET /v1/resource-manager/domains/{domainId}/provider/{service}/type/{resourceType}/resources
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", domainId, "provider", service, "type", resourceType, "resources").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ResPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractResources(pages)
}

type ResPage struct {
	pagination.NewSinglePageBase
}

func ExtractResources(r pagination.NewPage) ([]Resource, error) {
	var s struct {
		Resources []Resource `json:"resources"`
	}
	err := extract.Into(bytes.NewReader((r.(ResPage)).Body), &s)
	return s.Resources, err
}

type Resource struct {
	// Specifies the resource ID.
	ID string `json:"id"`
	// Specifies the resource name.
	Name string `json:"name"`
	// Specifies the cloud service name.
	Service string `json:"provider"`
	// Specifies the resource type.
	Type string `json:"type"`
	// Specifies the region ID.
	RegionId string `json:"region_id"`
	// Specifies the project ID in IaaS OpenStack.
	ProjectId string `json:"project_id"`
	// Specifies the project name in IaaS OpenStack.
	ProjectName string `json:"project_name"`
	// Specifies the resource checksum.
	Checksum string `json:"checksum"`
	// Specifies the time when the resource was created.
	CreatedAt string `json:"created"`
	// Specifies the time when the resource was updated.
	UpdatedAt string `json:"updated"`
	// Specifies the status of a resource operation.
	ProvisioningState string `json:"provisioning_state"`
	// Resource state. The value can be normal or deleted.
	State string `json:"state"`
	// Specifies the resource tag.
	Tags map[string]string `json:"tags"`
	// Specifies the detailed properties of the resource.
	Properties map[string]interface{} `json:"properties"`
}
