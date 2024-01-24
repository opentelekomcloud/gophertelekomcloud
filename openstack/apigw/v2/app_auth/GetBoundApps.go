package app_auth

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListAppsBoundOpts struct {
	GatewayID string `json:"-"`
	AppID     string `q:"app_id"`
	ApiID     string `q:"api_id"`
	ApiName   string `q:"api_name"`
	EnvID     string `q:"env_id"`
}

func ListAppsBound(client *golangsdk.ServiceClient, opts ListAppsBoundOpts) ([]ApiAuth, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "app-auths",
			"binded-apps") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BindingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractAuths(pages)
}
