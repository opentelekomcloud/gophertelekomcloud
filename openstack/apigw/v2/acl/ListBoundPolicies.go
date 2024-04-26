package acl

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListBindingOpts struct {
	GatewayID string `json:"-"`
	// Offset from which the query starts. If the value is less than 0, it is automatically converted to 0.
	Offset *int `q:"offset"`
	// Number of items displayed on each page.
	// A value less than or equal to 0 will be automatically converted to 20,
	// and a value greater than 500 will be automatically converted to 500.
	Limit int `q:"limit"`
	// API ID.
	ApiId string `q:"api_id" required:"true"`
	// Environment ID.
	EnvId string `q:"env_id"`
	// Environment name.
	EnvName string `q:"env_name"`
	// Access control policy ID.
	PolicyId string `q:"acl_id"`
	// Access control policy name.
	PolicyName string `q:"acl_name"`
}

func ListBoundPolicies(client *golangsdk.ServiceClient, opts ListBindingOpts) ([]ApiBindAclInfo, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("apigw", "instances", opts.GatewayID, "acl-bindings", "binded-acls").
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
	return ExtractPolicies(pages)
}

func ExtractPolicies(r pagination.NewPage) ([]ApiBindAclInfo, error) {
	var s struct {
		Bindings []ApiBindAclInfo `json:"acls"`
	}
	err := extract.Into(bytes.NewReader((r.(BindingPage)).Body), &s)
	return s.Bindings, err
}

type ApiBindAclInfo struct {
	// Access control policy ID.
	AclId string `json:"acl_id"`
	// Access control policy name.
	AclName string `json:"acl_name"`
	// Object type.
	// Enumeration values:
	// IP
	// DOMAIN
	// DOMAIN_ID
	EntityType string `json:"entity_type"`
	// Access control type.
	// PERMIT: whitelist
	// DENY: blacklist
	AclType string `json:"acl_type"`
	// Access control objects.
	AclValue string `json:"acl_value"`
	// Effective environment ID.
	EnvId string `json:"env_id"`
	// Effective environment name.
	EnvName string `json:"env_name"`
	// Binding record ID.
	BindID string `json:"bind_id"`
	// Binding time.
	BindTime string `json:"bind_time"`
}
