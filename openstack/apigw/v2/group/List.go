package group

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID string `json:"-"`
	// API group ID.
	ID string `q:"id"`
	// API group name.
	Name string `q:"name"`
	// Offset from which the query starts.
	// If the offset is less than 0, the value is automatically converted to 0. Default to 0.
	Offset int `q:"offset"`
	// Number of items displayed on each page. The valid values are range form 1 to 500, default to 20.
	Limit int `q:"limit"`
	// Parameter name for exact matching. Only API group names are supported.
	PreciseSearch string `q:"precise_search"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]GroupResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "api-groups") + q.String(),
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
