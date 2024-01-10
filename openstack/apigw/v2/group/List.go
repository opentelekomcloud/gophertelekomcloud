package group

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Limit        int    `q:"limit"`
	InstanceID   string `json:"-"`
	InstanceName string `q:"instance_name"`
	Status       string `q:"status"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]GroupResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.InstanceID, "api-groups") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return GroupPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractGateways(pages)
}

type GroupPage struct {
	pagination.NewSinglePageBase
}

func ExtractGateways(r pagination.NewPage) ([]GroupResp, error) {
	var s struct {
		Gateways []GroupResp `json:"groups"`
	}
	err := extract.Into(bytes.NewReader((r.(GroupPage)).Body), &s)
	return s.Gateways, err
}
