package compliance

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListRulesOpts struct {
	DomainId string `json:"-"`
	// Specifies whether resources are collected by default. tracked indicates that resources are collected by default,
	// and untracked indicates that resources are not collected by default.
	Track string `q:"policy_assignment_name"`
	// Specifies the number of records returned on each page during pagination query.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"string"`
}

func ListRules(client *golangsdk.ServiceClient, opts ListRulesOpts) ([]PolicyRule, error) {
	// GET /v1/resource-manager/domains/{domain_id}/policy-assignments
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", opts.DomainId, "policy-assignments").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ResPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractRules(pages)
}

func ExtractRules(r pagination.NewPage) ([]PolicyRule, error) {
	var s struct {
		Rules []PolicyRule `json:"value"`
	}
	err := extract.Into(bytes.NewReader((r.(ResPage)).Body), &s)
	return s.Rules, err
}
