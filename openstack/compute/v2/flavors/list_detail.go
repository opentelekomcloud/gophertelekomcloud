package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

/*
AccessType maps to OpenStack's Flavor.is_public field. Although the is_public
field is boolean, the request options are ternary, which is why AccessType is
a string. The following values are allowed:

The AccessType argument is optional, and if it is not supplied, OpenStack
returns the PublicAccess flavors.
*/
type AccessType string

const (
	// PublicAccess returns public flavors and private flavors associated with that project.
	PublicAccess AccessType = "true"
	// PrivateAccess (admin only) returns private flavors, across all projects.
	PrivateAccess AccessType = "false"
	// AllAccess (admin only) returns public and private flavors across all projects.
	AllAccess AccessType = "None"
)

/*
ListOpts filters the results returned by the List() function.
For example, a flavor with a minDisk field of 10 will not be returned if you
specify MinDisk set to 20.

Typically, software will use the last ID of the previous call to List to set
the Marker for the current call.
*/
type ListOpts struct {
	// ChangesSince, if provided, instructs List to return only those things which
	// have changed since the timestamp provided.
	ChangesSince string `q:"changes-since"`
	// MinDisk and MinRAM, if provided, elides flavors which do not meet your criteria.
	MinDisk int `q:"minDisk"`
	MinRAM  int `q:"minRam"`
	// SortDir allows to select sort direction. It can be "asc" or "desc" (default).
	SortDir string `q:"sort_dir"`
	// SortKey allows to sort by one of the flavors attributes. Default is flavorId.
	SortKey string `q:"sort_key"`
	// Marker and Limit control paging. Marker instructs List where to start listing from.
	Marker string `q:"marker"`
	// Limit instructs List to refrain from sending excessively large lists of flavors.
	Limit int `q:"limit"`
	// AccessType, if provided, instructs List which set of flavors to return.
	// If IsPublic not provided, flavors for the current project are returned.
	AccessType AccessType `q:"is_public"`
}

// ListDetail instructs OpenStack to provide a list of flavors.
// You may provide criteria by which List curtails its results for easier processing.
func ListDetail(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("flavors", "detail")+query.String(),
		func(r pagination.PageResult) pagination.Page {
			return FlavorPage{pagination.LinkedPageBase{PageResult: r}}
		})
}

// FlavorPage contains a single page of all flavors from a ListDetails call.
type FlavorPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines if a FlavorPage contains any results.
func (page FlavorPage) IsEmpty() (bool, error) {
	flavors, err := ExtractFlavors(page)
	return len(flavors) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the next page of results.
func (page FlavorPage) NextPageURL() (string, error) {
	var res []golangsdk.Link
	err := extract.IntoSlicePtr(page.BodyReader(), &res, "flavors_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res)
}

// ExtractFlavors provides access to the list of flavors in a page acquired from the ListDetail operation.
func ExtractFlavors(r pagination.Page) ([]Flavor, error) {
	var res []Flavor
	err := extract.IntoSlicePtr(r.(FlavorPage).BodyReader(), &res, "flavors")
	return res, err
}
