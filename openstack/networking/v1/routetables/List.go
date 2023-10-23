package routetables

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// RouteTablePage is the page returned by a pager when traversing over a
// collection of route tables
type RouteTablePage struct {
	pagination.NewSinglePageBase
}

// LastMarker returns the last route table ID in a ListResult
func (r RouteTablePage) LastMarker() (string, error) {
	tables, err := ExtractRouteTables(r)
	if err != nil {
		return "", err
	}
	if len(tables) == 0 {
		return "", nil
	}
	return tables[len(tables)-1].ID, nil
}

// IsEmpty checks whether a RouteTablePage struct is empty.
func (r RouteTablePage) IsEmpty() (bool, error) {
	tables, err := ExtractRouteTables(r)
	return len(tables) == 0, err
}

// ExtractRouteTables accepts a Page struct, specifically a RouteTablePage struct,
// and extracts the elements into a slice of RouteTable structs.
func ExtractRouteTables(r pagination.NewPage) ([]RouteTable, error) {
	var s struct {
		RouteTables []RouteTable `json:"routetables"`
	}
	err := extract.Into(bytes.NewReader((r.(RouteTablePage)).Body), &s)
	return s.RouteTables, err
}

// ListOpts allows to query all route tables or filter collections by parameters
// Marker and Limit are used for pagination.
type ListOpts struct {
	// ID is the unique identifier for the route table
	ID string `q:"id"`
	// VpcID is the unique identifier for the vpc
	VpcID string `q:"vpc_id"`
	// SubnetID the unique identifier for the subnet
	SubnetID string `q:"subnet_id"`
	// Limit is the number of records returned for each page query, the value range is 0~intmax
	Limit int `q:"limit"`
	// Marker is the starting resource ID of the paging query,
	// which means that the query starts from the next record of the specified resource
	Marker string `q:"marker"`
}

// List returns a Pager which allows you to iterate over a collection of
// vpc route tables. It accepts a ListOpts struct, which allows you to
// filter  the returned collection for greater efficiency.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]RouteTable, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(client.ProjectID, "routetables") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return RouteTablePage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	return ExtractRouteTables(pages)
}
