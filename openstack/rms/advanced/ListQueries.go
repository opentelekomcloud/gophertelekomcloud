package advanced

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListQueriesOpts struct {
	DomainId string `json:"-"`
	Name     string `q:"name"`
	// Specifies the number of records returned on each page during pagination query.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"string"`
}

func ListQueries(client *golangsdk.ServiceClient, opts ListQueriesOpts) ([]Query, error) {
	// GET /v1/resource-manager/domains/{domain_id}/stored-queries
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", opts.DomainId, "stored-queries").
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
	return ExtractQueries(pages)
}

type ResPage struct {
	pagination.NewSinglePageBase
}

func ExtractQueries(r pagination.NewPage) ([]Query, error) {
	var s struct {
		Queries []Query `json:"value"`
	}
	err := extract.Into(bytes.NewReader((r.(ResPage)).Body), &s)
	return s.Queries, err
}
