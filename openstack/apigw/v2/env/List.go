package env

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID string `json:"-"`
	Limit     int    `q:"limit"`
	Name      string `q:"name"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]EnvResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "envs") + q.String(),
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

func ExtractEnvs(r pagination.NewPage) ([]EnvResp, error) {
	var s struct {
		Gateways []EnvResp `json:"envs"`
	}
	err := extract.Into(bytes.NewReader((r.(EnvPage)).Body), &s)
	return s.Gateways, err
}
