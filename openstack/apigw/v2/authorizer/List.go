package authorizer

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID string `json:"-"`
	// Offset from which the query starts.
	// If the value is less than 0, it is automatically converted to 20.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 20,
	// and a value greater than 500 will be automatically converted to 500.
	Limit int `q:"limit"`
	// Authorizer id
	ID string `q:"id"`
	// Authorizer name
	Name string `q:"mame,"`
	// Authorizer type
	Type string `q:"type"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]AuthorizerResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", opts.GatewayID, "authorizers").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return AuthorizerPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractAuthorizers(pages)
}

type AuthorizerPage struct {
	pagination.NewSinglePageBase
}

func ExtractAuthorizers(r pagination.NewPage) ([]AuthorizerResp, error) {
	var s struct {
		Gateways []AuthorizerResp `json:"authorizer_list"`
	}
	err := extract.Into(bytes.NewReader((r.(AuthorizerPage)).Body), &s)
	return s.Gateways, err
}
