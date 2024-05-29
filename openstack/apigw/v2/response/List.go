package response

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows to filter list data using given parameters.
type ListOpts struct {
	// Offset from which the query starts.
	// If the offset is less than 0, the value is automatically converted to 0. Default to 0.
	Offset int `q:"offset"`
	// Number of items displayed on each page. The valid values are range form 1 to 500, default to 20.
	Limit int `q:"limit"`
	// Gateway ID.
	GatewayID string `json:"-"`
	// Group ID.
	GroupID string `json:"-"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Response, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", opts.GatewayID, "api-groups", opts.GroupID, "gateway-responses").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ResponsePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractResponses(pages)
}

type ResponsePage struct {
	pagination.NewSinglePageBase
}

func ExtractResponses(r pagination.NewPage) ([]Response, error) {
	var s struct {
		Responses []Response `json:"responses"`
	}
	err := extract.Into(bytes.NewReader((r.(ResponsePage)).Body), &s)
	return s.Responses, err
}
