package flavors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Flavor struct {
	// Specifies the ID of the flavor.
	ID string `json:"id"`

	// Specifies the info of the flavor.
	Info FlavorInfo `json:"info"`

	// Specifies the name of the flavor.
	Name string `json:"name"`

	// Specifies whether shared.
	Shared bool `json:"shared"`

	// Specifies the type of the flavor.
	Type string `json:"type"`

	// Specifies whether sold out.
	SoldOut bool `json:"flavor_sold_out"`
}

type FlavorInfo struct {
	// Specifies the connection
	Connection int `json:"connection"`

	// Specifies the cps.
	Cps int `json:"cps"`

	// Specifies the qps
	Qps int `json:"qps"`

	// Specifies the https_cps
	HttpsCps int `json:"https_cps"`

	// Specifies the lcu
	Lcu int `json:"lcu"`

	// Specifies the bandwidth
	Bandwidth int `json:"bandwidth"`
}

// FlavorPage is the page returned by a pager when traversing over a
// collection of flavor.
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
	var s []Flavor

	err := extract.IntoSlicePtr((r.(FlavorPage)).Body, &s, "flavors")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type GetResult struct {
	golangsdk.Result
}

func (r GetResult) Extract() (*Flavor, error) {
	s := new(Flavor)
	err := r.ExtractIntoStructPtr(s, "flavor")
	if err != nil {
		return nil, err
	}
	return s, nil
}
