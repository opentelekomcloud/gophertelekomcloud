package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts holds options for listing Volumes. It is passed to the volumes.List function.
type ListOpts struct {
	// AllTenants will retrieve volumes of all tenants/projects.
	AllTenants bool `q:"all_tenants"`
	// Metadata will filter results based on specified metadata.
	Metadata map[string]string `q:"metadata"`
	// Name will filter by the specified volume name.
	Name string `q:"name"`
	// Status will filter by the specified status.
	Status string `q:"status"`
	// TenantID will filter by a specific tenant/project ID.
	// Setting AllTenants is required for this.
	TenantID string `q:"project_id"`
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

// List returns Volumes optionally limited by the conditions provided in ListOpts.
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("volumes", "detail")+q.String(),
		func(r pagination.PageResult) pagination.Page {
			return VolumePage{pagination.LinkedPageBase{PageResult: r}}
		})
}

// VolumePage is a pagination.pager that is returned from a call to the List function.
type VolumePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a ListResult contains no Volumes.
func (r VolumePage) IsEmpty() (bool, error) {
	volumes, err := ExtractVolumes(r)
	return len(volumes) == 0, err
}

func (r VolumePage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"volumes_links"`
	}
	err := extract.Into(r.BodyReader(), &s)
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s.Links)
}

// ExtractVolumes extracts and returns Volumes. It is used while iterating over a volumes.List call.
func ExtractVolumes(r pagination.Page) ([]Volume, error) {
	var s []Volume
	err := ExtractVolumesInto(r, &s)
	return s, err
}

// ExtractVolumesInto similar to ExtractInto but operates on a `list` of volumes
func ExtractVolumesInto(r pagination.Page, v interface{}) error {
	return extract.IntoSlicePtr(r.(VolumePage).Result.BodyReader(), v, "volumes")
}
