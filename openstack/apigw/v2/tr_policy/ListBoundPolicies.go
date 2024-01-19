package throttling_policy

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListBindingOpts struct {
	GatewayID    string `json:"-"`
	ApiID        string `q:"api_id"`
	ThrottleID   string `q:"throttle_id"`
	ThrottleName string `q:"throttle_name"`
	EnvID        string `q:"env_id"`
}

func ListBoundPolicies(client *golangsdk.ServiceClient, opts ListBindingOpts) ([]ThrottlingBindResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "throttle-bindings",
			"binded-throttles") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BindingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractPolicies(pages)
}

func ExtractPolicies(r pagination.NewPage) ([]ThrottlingBindResp, error) {
	var s struct {
		Bindings []ThrottlingBindResp `json:"throttles"`
	}
	err := extract.Into(bytes.NewReader((r.(BindingPage)).Body), &s)
	return s.Bindings, err
}

type ThrottlingBindResp struct {
	AppCallLimits          int    `json:"app_call_limits"`
	Name                   string `json:"name"`
	TimeUnit               string `json:"time_unit"`
	Description            string `json:"remark"`
	ApiCallLimits          int    `json:"api_call_limits"`
	Type                   int    `json:"type"`
	EnableAdaptiveControl  string `json:"enable_adaptive_control"`
	UserCallLimits         int    `json:"user_call_limits"`
	TimeInterval           int    `json:"time_interval"`
	IpCallLimits           int    `json:"ip_call_limits"`
	ID                     string `json:"id"`
	BindNum                int    `json:"bind_num"`
	IsIncluSpecialThrottle int    `json:"is_inclu_special_throttle"`
	CreateTime             string `json:"create_time"`
	EnvName                string `json:"env_name"`
	BindID                 string `json:"bind_id"`
	BindTime               string `json:"bind_time"`
}
