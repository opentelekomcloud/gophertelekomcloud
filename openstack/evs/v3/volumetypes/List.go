package volumetypes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts holds options for listing Volume Types. It is passed to the volumetypes.List function.
type ListOpts struct {
	// Comma-separated list of sort keys and optional sort directions in the
	// form of <key>[:<direction>].
	Sort string `q:"sort"`
	// Requests a page size of items.
	Limit int `q:"limit"`
	// Used in conjunction with limit to return a slice of items.
	Offset int `q:"offset"`
	// The ID of the last-seen item.
	Marker string `q:"marker"`
}

// ToVolumeTypeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeTypeListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// List returns Volume types.
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("types")+q.String(), func(r pagination.PageResult) pagination.Page {
		return VolumeTypePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// VolumeTypePage is a pagination.pager that is returned from a call to the List function.
type VolumeTypePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a ListResult contains no Volume Types.
func (r VolumeTypePage) IsEmpty() (bool, error) {
	volumeTypes, err := ExtractVolumeTypes(r)
	return len(volumeTypes) == 0, err
}

func (r VolumeTypePage) NextPageURL() (string, error) {
	var s []golangsdk.Link
	err := extract.IntoSlicePtr(r.BodyReader(), &s, "volume_type_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s)
}

// ExtractVolumeTypesInto similar to ExtractInto but operates on a `list` of volume types
func ExtractVolumeTypesInto(r pagination.Page, v interface{}) error {
	return extract.IntoSlicePtr(r.(VolumeTypePage).BodyReader(), v, "volume_types")
}

// ExtractVolumeTypes extracts and returns Volumes. It is used while iterating over a volumetypes.List call.
func ExtractVolumeTypes(r pagination.Page) ([]VolumeType, error) {
	var s []VolumeType
	err := ExtractVolumeTypesInto(r, &s)
	return s, err
}
