package shares

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Limit  string `q:"limit"`
	Offset string `q:"offset"`
}

// List returns a Pager which allows you to iterate over a collection of
// SFS Turbo resources.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Turbo, error) {
	// GET /v1/{project_id}/sfs-turbo/shares/detail
	url, err := golangsdk.NewURLBuilder().WithEndpoints("sfs-turbo", "shares", "detail").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SharePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractInstances(pages)
}

func ExtractInstances(r pagination.NewPage) ([]Turbo, error) {
	var ListResponse ListResponse
	err := extract.Into(bytes.NewReader((r.(SharePage)).Body), &ListResponse)
	return ListResponse.Shares, err
}

type SharePage struct {
	pagination.NewSinglePageBase
}

type ListResponse struct {
	// List of SFS Turbo file systems
	Shares []Turbo `json:"shares"`
	// Number of SFS Turbo file systems
	Count int `json:"count"`
}
