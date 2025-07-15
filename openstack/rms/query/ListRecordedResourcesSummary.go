package query

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListRecordedSummaryOpts struct {
	// Specifies the resource name.
	Name string `q:"name"`
	// Specifies the resource type
	Type string `q:"type"`
	// Specifies the region ID.
	RegionId string `q:"region_id"`
	// Specifies the project ID.
	ProjectId string `q:"project_id"`
	// Specifies tags. The format is key or key=value.
	Tags []string `q:"tags"`
	// Indicating whether deleted resources need to be returned.
	// This parameter is set to false by default.
	ResourceDeleted *bool `q:"resource_deleted"`
}

func ListRecordedResourcesSummary(client *golangsdk.ServiceClient, domainId string, opts ListRecordedSummaryOpts) ([]Summary, error) {
	// GET /v1/resource-manager/domains/{domain_id}/tracked-resources/summary
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", domainId, "tracked-resources", "summary").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SummaryPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractSummary(pages)
}

type SummaryPage struct {
	pagination.NewSinglePageBase
}

func ExtractSummary(r pagination.NewPage) ([]Summary, error) {
	var summaries []Summary
	err := extract.Into(bytes.NewReader((r.(SummaryPage)).Body), &summaries)
	return summaries, err
}

type Summary struct {
	// Specifies the cloud service name.
	Provider string `json:"provider"`
	// Specifies the resource type list.
	Types []Types `json:"types"`
}

type Types struct {
	// Specifies the resource type.
	Type string `json:"type"`
	// Specifies the regions.
	Regions []Regions `json:"regions"`
}

type Regions struct {
	// Specifies the region ID.
	RegionId string `json:"region_id"`
	// Specifies the number of resources of this type in the current region.
	Count int `json:"count"`
}
