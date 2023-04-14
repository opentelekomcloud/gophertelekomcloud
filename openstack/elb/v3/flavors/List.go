package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through the API.
type ListOpts struct {
	// Specifies the id.
	ID []string `q:"id"`
	// Specifies the name.
	Name []string `q:"name"`
	// Specifies whether the flavor is available to all users.
	//
	// true indicates that the flavor is available to all users.
	//
	// false indicates that the flavor is available only to a specific user.
	Shared *bool `q:"shared"`
	// Specifies the type.
	Type []string `q:"type"`
	// Specifies the ID of the last record on the previous page.
	//
	// Note:
	//
	// This parameter must be used together with limit.
	//
	// If this parameter is not specified, the first page will be queried.
	//
	// This parameter cannot be left blank or set to an invalid ID.
	Marker string `q:"marker"`
	// Specifies the number of records on each page.
	//
	// Minimum: 0
	//
	// Maximum: 2000
	Limit *int `q:"limit"`
	// Specifies the page direction.
	//
	// The value can be true or false, and the default value is false.
	//
	// The last page in the list requested with page_reverse set to false will not contain the "next" link, and the last page in the list requested with page_reverse set to true will not contain the "previous" link.
	//
	// This parameter must be used together with limit.
	PageReverse *bool `q:"page_reverse"`
}

// List returns a Pager which allows you to iterate over a collection of flavors.
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	queryString, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("flavors")+queryString.String(), func(r pagination.PageResult) pagination.Page {
		return FlavorPage{PageWithInfo: pagination.NewPageWithInfo(r)}
	})
}

// FlavorPage is the page returned by a pager when traversing over a collection of flavor.
type FlavorPage struct {
	pagination.PageWithInfo
}

// IsEmpty checks whether a FlavorsPage struct is empty.
func (r FlavorPage) IsEmpty() (bool, error) {
	is, err := ExtractFlavors(r)
	return len(is) == 0, err
}

// ExtractFlavors accepts a Page struct, specifically a FlavorsPage struct,
// and extracts the elements into a slice of flavor structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractFlavors(r pagination.Page) ([]Flavor, error) {
	var res []Flavor
	err := extract.IntoSlicePtr(r.(FlavorPage).BodyReader(), &res, "flavors")
	return res, err
}
