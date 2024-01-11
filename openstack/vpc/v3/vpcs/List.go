package vpcs

import (
	"bytes"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// Number of records displayed on each page.
	// Value range: 0 to 2000
	Limit int `q:"limit,omitempty"`
	// Start resource ID of pagination query.
	// If the parameter is left blank, only resources on the first page are queried.
	Marker string `q:"marker,omitempty"`
	// VPC ID, which can be used to filter VPCs.
	Id []string `q:"id,omitempty"`
	// VPC name, which can be used to filter VPCs.
	Name []string `q:"name,omitempty"`
	// Supplementary information about the VPC, which can be used to filter VPCs.
	Description []string `q:"description,omitempty"`
	// VPC CIDR block, which can be used to filter VPCs.
	Cidr []string `q:"cidr,omitempty"`
}

// List is used to query VPCs.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Vpc, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("vpcs") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return VpcPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractVpcs(pages)
}

// VpcPage is the page returned by a pager when traversing over a
// collection of vpcs
type VpcPage struct {
	pagination.NewSinglePageBase
}

// ExtractVpcs accepts a Page struct, specifically a VpcPage struct,
// and extracts the elements into a slice of Vpc structs.
func ExtractVpcs(r pagination.NewPage) ([]Vpc, error) {
	var s struct {
		Vpcs []Vpc `json:"vpcs"`
	}
	err := extract.Into(bytes.NewReader((r.(VpcPage)).Body), &s)
	return s.Vpcs, err
}
