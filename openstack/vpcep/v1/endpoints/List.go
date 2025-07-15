package endpoints

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Specifies the name of the VPC endpoint service. The value is not case-sensitive and supports fuzzy match.
	Name string `q:"endpoint_service_name"`
	// Specifies the ID of the VPC where the VPC endpoint is to be created.
	VpcID string `q:"vpc_id"`
	// Specifies the unique ID of the VPC endpoint.
	ID string `q:"id"`
	// Specifies the sorting field of the VPC endpoint list. The field can be:
	// created_at: VPC endpoints are sorted by creation time.
	// updated_at: VPC endpoints are sorted by update time.
	// The default field is created_at.
	SortKey string `q:"sort_key"`
	// Specifies the sorting method of the VPC endpoint list. The method can be:
	// desc: VPC endpoints are sorted in descending order.
	// asc: VPC endpoints are sorted in ascending order.
	// The default method is desc.
	SortDir string `q:"sort_dir"`
	// Specifies the maximum number of VPC endpoint services displayed on each page.
	Limit *int `q:"limit"`
	// Specifies the offset.
	Offset *int `q:"offset"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Endpoint, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("vpc-endpoints").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return EndpointPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractEndpoints(pages)
}

type EndpointPage struct {
	pagination.NewSinglePageBase
}

func ExtractEndpoints(r pagination.NewPage) ([]Endpoint, error) {
	var s struct {
		Endpoints []Endpoint `json:"endpoints"`
	}
	err := extract.Into(bytes.NewReader((r.(EndpointPage)).Body), &s)
	return s.Endpoints, err
}
