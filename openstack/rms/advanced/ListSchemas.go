package advanced

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListSchemasOpts struct {
	DomainId string `json:"-"`
	Name     string `q:"name"`
	// Specifies the number of records returned on each page during pagination query.
	Limit *int `q:"limit"`
	// Specifies the pagination parameter.
	Marker string `q:"string"`
}

func ListSchemas(client *golangsdk.ServiceClient, opts ListSchemasOpts) ([]Schema, error) {
	// GET /v1/resource-manager/domains/{domain_id}/schemas
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("resource-manager", "domains", opts.DomainId, "schemas").
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
	return ExtractSchemas(pages)
}

func ExtractSchemas(r pagination.NewPage) ([]Schema, error) {
	var s struct {
		Schemas []Schema `json:"value"`
	}
	err := extract.Into(bytes.NewReader((r.(ResPage)).Body), &s)
	return s.Schemas, err
}

type Schema struct {
	Type   string      `json:"type"`
	Schema interface{} `json:"schema"`
}
