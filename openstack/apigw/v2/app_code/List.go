package app_code

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID string `json:"-"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]CodeResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "apps") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return AppCodePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractAppCodes(pages)
}

type AppCodePage struct {
	pagination.NewSinglePageBase
}

func ExtractAppCodes(r pagination.NewPage) ([]CodeResp, error) {
	var s struct {
		AppCodes []CodeResp `json:"app_codes"`
	}
	err := extract.Into(bytes.NewReader((r.(AppCodePage)).Body), &s)
	return s.AppCodes, err
}
