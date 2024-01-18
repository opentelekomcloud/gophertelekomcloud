package key

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID     string `json:"-"`
	SignatureID   string `q:"id"`
	SignatureName string `q:"name"`
	PreciseSearch string `q:"precise_search"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]SignKeyResp, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("apigw", "instances", opts.GatewayID, "signs") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SignatureKeyPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractSignatureKey(pages)
}

type SignatureKeyPage struct {
	pagination.NewSinglePageBase
}

func ExtractSignatureKey(r pagination.NewPage) ([]SignKeyResp, error) {
	var s struct {
		SignatureKeys []SignKeyResp `json:"signs"`
	}
	err := extract.Into(bytes.NewReader((r.(SignatureKeyPage)).Body), &s)
	return s.SignatureKeys, err
}
