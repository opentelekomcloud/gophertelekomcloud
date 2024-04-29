package acl

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListBoundOpts struct {
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

func ListAPIBoundPolicy(client *golangsdk.ServiceClient, opts ListBoundOpts) ([]ApiAcl, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", opts.GatewayID, "acl-bindings", "binded-apis").
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

type BindingPage struct {
	pagination.NewSinglePageBase
}

func ExtractBindings(r pagination.NewPage) ([]ApiAcl, error) {
	var s struct {
		Bindings []ApiAcl `json:"apis"`
	}
	err := extract.Into(bytes.NewReader((r.(BindingPage)).Body), &s)
	return s.Bindings, err
}

type ApiAcl struct {
	// API ID.
	ApiId string `json:"api_id"`
	// API name.
	ApiName string `json:"api_name"`
	// API type.
	ApiType int `json:"api_type"`
	// API description.
	ApiDescription string `json:"api_remark"`
	// ID of the environment in which the policy takes effect.
	EnvId string `json:"env_id"`
	// Name of the environment in which the policy takes effect.
	EnvName string `json:"env_name"`
	// Binding record ID.
	BindingId string `json:"bind_id"`
	// API group name.
	GroupName string `json:"group_name"`
	// Binding time.
	BindedAt string `json:"bind_time"`
	// API publication record ID.
	PublishId string `json:"publish_id"`
	// Request method.
	ReqMethod string `json:"req_method"`
}
