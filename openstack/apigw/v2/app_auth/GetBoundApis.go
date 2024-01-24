package app_auth

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListBoundOpts struct {
	GatewayID string `json:"-"`
	AppID     string `q:"app_id" required:"true"`
	ApiID     string `q:"api_id"`
	ApiName   string `q:"api_name"`
	GroupID   string `q:"group_id"`
	GroupName string `q:"group_name"`
	EnvID     string `q:"env_id"`
}

func ListAPIBound(client *golangsdk.ServiceClient, opts ListBoundOpts) ([]ApiAuth, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "app-auths",
			"binded-apis") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BindingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractAuths(pages)
}

type BindingPage struct {
	pagination.NewSinglePageBase
}

func ExtractAuths(r pagination.NewPage) ([]ApiAuth, error) {
	var s struct {
		Bindings []ApiAuth `json:"auths"`
	}
	err := extract.Into(bytes.NewReader((r.(BindingPage)).Body), &s)
	return s.Bindings, err
}

type ApiAuth struct {
	ID             string   `json:"id"`
	ApiID          string   `json:"api_id"`
	ApiName        string   `json:"api_name"`
	GroupName      string   `json:"group_name"`
	ApiType        int      `json:"api_type"`
	ApiDescription string   `json:"api_remark"`
	EnvID          string   `json:"env_id"`
	AuthRole       string   `json:"auth_role"`
	AuthTime       string   `json:"auth_time"`
	AppName        string   `json:"app_name"`
	AppDescription string   `json:"app_remark"`
	AppType        string   `json:"app_type"`
	AppCreator     string   `json:"app_creator"`
	PublishID      string   `json:"publish_id"`
	GroupID        string   `json:"group_id"`
	AuthTunnel     string   `json:"auth_tunnel"`
	AuthWhitelist  []string `json:"auth_whitelist"`
	AuthBlacklist  []string `json:"auth_blacklist"`
	VisitParam     string   `json:"visit_param"`
	EnvName        string   `json:"env_name"`
	AppID          string   `json:"app_id"`
}
