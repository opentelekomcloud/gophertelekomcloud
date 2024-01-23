package special_policy

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID  string `json:"-"`
	ThrottleID string `json:"-"`
	ObjectType string `q:"object_type"`
	AppName    string `q:"app_name"`
	User       string `q:"user"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]ThrottlingResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client: client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "throttles",
			opts.ThrottleID, "throttle-specials") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ThrottlingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractTrPolicy(pages)
}

type ThrottlingPage struct {
	pagination.NewSinglePageBase
}

func ExtractTrPolicy(r pagination.NewPage) ([]ThrottlingResp, error) {
	var s struct {
		ThrottlingPolicies []ThrottlingResp `json:"throttle_specials"`
	}
	err := extract.Into(bytes.NewReader((r.(ThrottlingPage)).Body), &s)
	return s.ThrottlingPolicies, err
}
