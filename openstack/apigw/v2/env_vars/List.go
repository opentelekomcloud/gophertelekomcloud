package env_vars

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID     string `json:"-"`
	GroupID       string `q:"group_id" required:"true"`
	EnvID         string `q:"env_id"`
	VariableName  string `q:"variable_name"`
	PreciseSearch string `q:"precise_search"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]EnvVarsResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "env-variables") + q.String(),
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

func ExtractEnvs(r pagination.NewPage) ([]EnvVarsResp, error) {
	var s struct {
		Gateways []EnvVarsResp `json:"variables"`
	}
	err := extract.Into(bytes.NewReader((r.(EnvPage)).Body), &s)
	return s.Gateways, err
}
