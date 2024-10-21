package customer_gateway

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the number of records returned on each page during pagination query.
	Limit *int `q:"limit"`
	// Specifies the start flag for querying the current page. If this parameter is left blank, the first page is queried.
	// The marker for querying the next page is the next_marker in the page_info object returned on the current page.
	Marker string `q:"marker"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]CustomerGateway, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("customer-gateways").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return GwPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractGateways(pages)
}

type GwPage struct {
	pagination.NewSinglePageBase
}

func ExtractGateways(r pagination.NewPage) ([]CustomerGateway, error) {
	var s struct {
		Gateways []CustomerGateway `json:"customer_gateways"`
	}
	err := extract.Into(bytes.NewReader((r.(GwPage)).Body), &s)
	return s.Gateways, err
}
