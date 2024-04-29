package acl

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListUnoundOpts struct {
	GatewayID string `json:"-"`
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 20,
	// and a value greater than 500 will be automatically converted to 500.
	Limit int `q:"limit"`
	// Access control policy ID.
	ID string `q:"acl_id" required:"true"`
	// API ID.
	ApiId string `q:"api_id"`
	// API name.
	ApiName string `q:"api_name"`
	// Environment ID.
	EnvId string `q:"env_id"`
	// API group ID.
	GroupId string `q:"group_id"`
}

func ListAPIUnoundPolicy(client *golangsdk.ServiceClient, opts ListBoundOpts) ([]ApiAcl, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", opts.GatewayID, "acl-bindings", "unbinded-apis").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return BindingPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractBindings(pages)
}
