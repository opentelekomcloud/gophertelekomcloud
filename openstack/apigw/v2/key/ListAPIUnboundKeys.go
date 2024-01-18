package key

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListUnbindOpts struct {
	GatewayID string `json:"-"`
	ApiID     string `q:"api_id"`
	SignID    string `q:"sign_id"`
	SignName  string `q:"sign_name"`
	EnvID     string `q:"env_id"`
}

func ListUnboundKeys(client *golangsdk.ServiceClient, opts ListUnbindOpts) ([]ApiForSign, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "sign-bindings",
			"unbinded-apis") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BindingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractApis(pages)
}

func ExtractApis(r pagination.NewPage) ([]ApiForSign, error) {
	var s struct {
		Apis []ApiForSign `json:"apis"`
	}
	err := extract.Into(bytes.NewReader((r.(BindingPage)).Body), &s)
	return s.Apis, err
}

type ApiForSign struct {
	AuthType      string `json:"auth_type"`
	RunEnvName    string `json:"run_env_name"`
	GroupName     string `json:"group_name"`
	PublishID     string `json:"publish_id"`
	GroupID       string `json:"group_id"`
	Name          string `json:"name"`
	Description   string `json:"Remark"`
	RunEnvID      string `json:"run_env_id"`
	ID            string `json:"id"`
	ReqURI        string `json:"req_uri"`
	Type          int    `json:"type"`
	SignatureName string `json:"signature_name"`
}
