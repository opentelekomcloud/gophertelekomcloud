package acl

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	GatewayID string `json:"-"`
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 20,
	// and a value greater than 500 will be automatically converted to 500.
	Limit int `q:"limit"`
	// Access control policy ID.
	ID string `q:"id"`
	// Access control policy name.
	Name string `q:"name"`
	// Type.
	// PERMIT (whitelist)
	// DENY (blacklist)
	Type string `q:"acl_type"`
	// Object types.
	// IP
	// DOMAIN
	EntityType string `q:"entity_type"`
	// Parameter name (name) for exact matching.
	PreciseSearch string `q:"precise_search"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]AclResp, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", opts.GatewayID, "acls").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return AclPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractAcls(pages)
}

type AclPage struct {
	pagination.NewSinglePageBase
}

func ExtractAcls(r pagination.NewPage) ([]AclResp, error) {
	var s struct {
		Acls []AclResp `json:"acls"`
	}
	err := extract.Into(bytes.NewReader((r.(AclPage)).Body), &s)
	return s.Acls, err
}
