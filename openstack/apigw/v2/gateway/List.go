package gateway

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Limit        int    `q:"limit"`
	InstanceID   string `q:"instance_id"`
	InstanceName string `q:"instance_name"`
	Status       string `q:"status"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Gateway, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return GatewayPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractGateways(pages)
}

type GatewayPage struct {
	pagination.NewSinglePageBase
}

func ExtractGateways(r pagination.NewPage) ([]Gateway, error) {
	var s struct {
		Gateways []Gateway `json:"instances"`
	}
	err := extract.Into(bytes.NewReader((r.(GatewayPage)).Body), &s)
	return s.Gateways, err
}
