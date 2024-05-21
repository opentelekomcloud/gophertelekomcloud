package app_code

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListAppsOpts struct {
	GatewayID string `json:"-"`
	AppID     string `json:"-"`
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 20,
	// and a value greater than 500 will be automatically converted to 500.
	Limit int `q:"limit"`
}

func ListAppCodesOfApp(client *golangsdk.ServiceClient, opts ListAppsOpts) ([]CodeResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", opts.GatewayID, "apps", opts.AppID, "app-codes").
		WithQueryParams(&opts).Build()
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return AppCodesOfAppPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractAppCodesOfApp(pages)
}

type AppCodesOfAppPage struct {
	pagination.NewSinglePageBase
}

func ExtractAppCodesOfApp(r pagination.NewPage) ([]CodeResp, error) {
	var s struct {
		AppCodes []CodeResp `json:"app_codes"`
	}
	err := extract.Into(bytes.NewReader((r.(AppCodesOfAppPage)).Body), &s)
	return s.AppCodes, err
}
