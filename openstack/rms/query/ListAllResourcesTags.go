package query

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListAllTagsOpts struct {
	// Specifies the maximum number of resources to return.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"marker"`
	// Specifies the name of the tag key.
	Key string `q:"key"`
}

func ListAllResourcesTags(client *golangsdk.ServiceClient, domainId string, opts ListAllTagsOpts) ([]Tag, error) {
	// GET /v1/resource-manager/domains/{domain_id}/all-resources/tags
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", domainId, "all-resources", "tags").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return TagsPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractTags(pages)
}
