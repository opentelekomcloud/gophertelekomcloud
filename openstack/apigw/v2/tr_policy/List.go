package throttling_policy

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID     string `json:"-"`
	ThrottleID    string `q:"id"`
	PolicyName    string `q:"name"`
	PreciseSearch string `q:"precise_search"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]ThrottlingResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "throttles") + q.String(),
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
		ThrottlingPolicies []ThrottlingResp `json:"throttles"`
	}
	err := extract.Into(bytes.NewReader((r.(ThrottlingPage)).Body), &s)
	return s.ThrottlingPolicies, err
}
