package app_auth

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListUnboundOpts struct {
	GatewayID string `json:"-"`
	AppID     string `q:"app_id" required:"true"`
	EnvID     string `q:"env_id" required:"true"`
	GroupID   string `q:"group_id"`
	ApiID     string `q:"api_id"`
	ApiName   string `q:"api_name"`
}

func ListAPIUnBound(client *golangsdk.ServiceClient, opts ListUnboundOpts) ([]ApiOutline, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "app-auths",
			"unbinded-apis") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BindingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractApiOutline(pages)
}

func ExtractApiOutline(r pagination.NewPage) ([]ApiOutline, error) {
	var s struct {
		Bindings []ApiOutline `json:"apis"`
	}
	err := extract.Into(bytes.NewReader((r.(BindingPage)).Body), &s)
	return s.Bindings, err
}

type ApiOutline struct {
	AuthType    string `json:"auth_type"`
	RunEnvName  string `json:"run_env_name"`
	GroupName   string `json:"group_name"`
	PublishID   string `json:"publish_id"`
	GroupID     string `json:"group_id"`
	Name        string `json:"name"`
	Description string `json:"remark"`
	RunEnvID    string `json:"run_env_id"`
	ID          string `json:"id"`
	ReqUri      string `json:"req_uri"`
}
