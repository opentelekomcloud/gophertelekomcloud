package key

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListBindingOpts struct {
	GatewayID string `json:"-"`
	ApiID     string `q:"api_id"`
	SignID    string `q:"sign_id"`
	SignName  string `q:"sign_name"`
	EnvID     string `q:"env_id"`
}

func ListBoundKeys(client *golangsdk.ServiceClient, opts ListBindingOpts) ([]BindSignResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "sign-bindings",
			"binded-signs") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BindingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractBindings(pages)
}

type BindingPage struct {
	pagination.NewSinglePageBase
}

func ExtractBindings(r pagination.NewPage) ([]BindSignResp, error) {
	var s struct {
		Bindings []BindSignResp `json:"bindings"`
	}
	err := extract.Into(bytes.NewReader((r.(BindingPage)).Body), &s)
	return s.Bindings, err
}
