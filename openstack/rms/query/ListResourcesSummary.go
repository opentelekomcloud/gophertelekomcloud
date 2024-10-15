package query

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListSummaryOpts struct {
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
}

func ListResourcesSummary(client *golangsdk.ServiceClient, domainId string, opts ListSummaryOpts) ([]Summary, error) {
	// GET /v1/resource-manager/domains/{domain_id}/all-resources/summary
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", domainId, "all-resources", "summary").
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
