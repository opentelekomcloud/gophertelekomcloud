package gateway

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListFeaturesOpts struct {
	GatewayID string `json:"-"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 20,
	// and a value greater than 500 will be automatically converted to 500.
	Limit int `q:"limit"`
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
}

func ListGatewayFeatures(client *golangsdk.ServiceClient, opts ListFeaturesOpts) ([]FeatureResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "features") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return FeaturePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractFeatures(pages)
}

type FeaturePage struct {
	pagination.NewSinglePageBase
}

func ExtractFeatures(r pagination.NewPage) ([]FeatureResp, error) {
	var s struct {
		Features []FeatureResp `json:"features"`
	}
	err := extract.Into(bytes.NewReader((r.(FeaturePage)).Body), &s)
	return s.Features, err
}
