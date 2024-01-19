package app

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID     string `json:"-"`
	ID            string `q:"id"`
	Name          string `q:"mame,"`
	Status        string `q:"status"`
	AppKey        string `q:"app_key"`
	Creator       string `q:"creator"`
	PreciseSearch string `q:"precise_search"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]AppResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "apps") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return EnvPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractEnvs(pages)
}

type EnvPage struct {
	pagination.NewSinglePageBase
}

func ExtractEnvs(r pagination.NewPage) ([]AppResp, error) {
	var s struct {
		Gateways []AppResp `json:"apps"`
	}
	err := extract.Into(bytes.NewReader((r.(EnvPage)).Body), &s)
	return s.Gateways, err
}
