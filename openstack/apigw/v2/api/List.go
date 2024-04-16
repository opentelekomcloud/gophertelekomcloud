package api

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID     string `json:"-"`
	ID            string `json:"id"`
	GroupID       string `q:"group_id"`
	ReqProtocol   string `q:"req_protocol"`
	ReqMethod     string `q:"req_method"`
	ReqUri        string `q:"req_uri"`
	AuthType      string `q:"auth_type"`
	PreciseSearch string `q:"precise_search"`
	EnvID         string `q:"env_id"`
	Type          int    `q:"type"`
	Name          string `q:"name"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]ApiResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "apis") + q.String(),
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

func ExtractEnvs(r pagination.NewPage) ([]ApiResp, error) {
	var s struct {
		Gateways []ApiResp `json:"apis"`
	}
	err := extract.Into(bytes.NewReader((r.(EnvPage)).Body), &s)
	return s.Gateways, err
}
