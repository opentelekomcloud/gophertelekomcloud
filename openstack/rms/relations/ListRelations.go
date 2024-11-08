package relations

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListAllOpts struct {
	DomainId   string `json:"-"`
	ResourceId string `json:"-"`
	// Specifies the direction of a resource relationship.
	Direction string `q:"direction"`
	// Specifies the maximum number of resources to return.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"marker"`
	// Specifies the region ID.
}

func ListRelations(client *golangsdk.ServiceClient, opts ListAllOpts) ([]ResourceRelation, error) {
	// GET /v1/resource-manager/domains/{domain_id}/all-resources/{resource_id}/relations
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", opts.DomainId, "all-resources", opts.ResourceId, "relations").
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

func ExtractResources(r pagination.NewPage) ([]ResourceRelation, error) {
	var s struct {
		Resources []ResourceRelation `json:"relations"`
	}
	err := extract.Into(bytes.NewReader((r.(ResPage)).Body), &s)
	return s.Resources, err
}

type ResourceRelation struct {
	RelationType     string `json:"relation_type"`
	FromResourceType string `json:"from_resource_type"`
	ToResourceType   string `json:"to_resource_type"`
	FromResourceId   string `json:"from_resource_id"`
	ToResourceId     string `json:"to_resource_id"`
}
