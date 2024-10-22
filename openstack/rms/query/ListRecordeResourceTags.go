package query

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListRecordedTagsOpts struct {
	// Specifies the maximum number of resources to return.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"marker"`
	// Specifies the name of the tag key.
	Key string `q:"key"`
	// Indicating whether deleted resources need to be returned.
	// This parameter is set to false by default.
	ResourceDeleted *bool `q:"resource_deleted"`
}

func ListRecordedResourcesTags(client *golangsdk.ServiceClient, domainId string, opts ListRecordedTagsOpts) ([]Tag, error) {
	// GET /v1/resource-manager/domains/{domain_id}/tracked-resources/tags
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", domainId, "tracked-resources", "tags").
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

type TagsPage struct {
	pagination.NewSinglePageBase
}

func ExtractTags(r pagination.NewPage) ([]Tag, error) {
	var s struct {
		Tags []Tag `json:"tags"`
	}
	err := extract.Into(bytes.NewReader((r.(TagsPage)).Body), &s)
	return s.Tags, err
}

type Tag struct {
	// Specifies the tag key.
	Key string `json:"key"`
	// Specifies tag values.
	Value []string `json:"value"`
}
