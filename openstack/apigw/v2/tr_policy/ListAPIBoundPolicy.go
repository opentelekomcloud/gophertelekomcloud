package throttling_policy

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListBoundOpts struct {
	GatewayID  string `json:"-"`
	ThrottleID string `q:"throttle_id"`
	EnvID      string `q:"env_id"`
	GroupID    string `q:"group_id"`
	ApiID      string `q:"api_id"`
	ApiName    string `q:"api_name"`
}

func ListAPIBoundPolicy(client *golangsdk.ServiceClient, opts ListBoundOpts) ([]ApiThrottle, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "throttle-bindings",
			"binded-apis") + q.String(),
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

func ExtractBindings(r pagination.NewPage) ([]ApiThrottle, error) {
	var s struct {
		Bindings []ApiThrottle `json:"apis"`
	}
	err := extract.Into(bytes.NewReader((r.(BindingPage)).Body), &s)
	return s.Bindings, err
}

type ApiThrottle struct {
	AuthType        string `json:"auth_type"`
	GroupName       string `json:"group_name"`
	PublishID       string `json:"publish_id"`
	ThrottleApplyID string `json:"throttle_apply_id"`
	ApplyTime       string `json:"apply_time"`
	Description     string `json:"remark"`
	RunEnvID        string `json:"run_env_id"`
	Type            int    `json:"int"`
	ThrottleName    string `json:"throttle_name"`
	ReqUri          string `json:"req_uri"`
	RunEnvName      string `json:"run_env_name"`
	GroupID         string `json:"group_id"`
	Name            string `json:"name"`
	ID              string `json:"id"`
	ReqID           string `json:"req_id"`
	ReqMethod       string `json:"req_method"`
}
