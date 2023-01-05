package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts holds options for listing Volumes. It is passed to the volumes.List function.
type ListOpts struct {
	// Specifies the AZ.
	AvailabilityZone bool `q:"availability_zone"`
	// Metadata will filter results based on specified metadata.
	Metadata map[string]string `q:"metadata"`
	// Specifies the disk name. The value can contain a maximum of 255 bytes.
	Name string `q:"name"`
	// Status will filter by the specified status.
	Status string `q:"status"`
	// TenantID will filter by a specific tenant/project ID.
	// Setting AllTenants is required for this.
	TenantID string `q:"project_id"`
	// Specifies the keyword based on which the returned results are sorted.
	// The value can be id, status, size, or created_at, and the default value is created_at.
	Sort string `q:"sort_key"`
	// Specifies the result sorting order. The default value is desc.
	// desc: specifies the descending order.
	// asc: specifies the ascending order.
	SortDir string `q:"sort_dir"`
	// Specifies the maximum number of query results that can be returned.
	// The value ranges from 1 to 1000, and the default value is 1000. The returned value cannot exceed this limit.
	// If the tenant has more than 50 disks in total, you are advised to use this parameter and set its value to 50 to improve the query efficiency.
	// Examples are provided as follows:
	// GET /v3/xxx/volumes?limit=50: Queries the 1–50 disks. GET /v3/xxx/volumes?offset=50&limit=50: Queries the 51–100 disks.
	Limit int `q:"limit"`
	// Specifies the offset.
	// All disks after this offset will be queried. The value must be an integer greater than 0 but less than the number of disks.
	Offset int `q:"offset"`
	// Specifies the ID of the last record on the previous page. The returned value is the value of the item after this one.
	Marker string `q:"marker"`
}

// List returns Volumes optionally limited by the conditions provided in ListOpts.
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	// GET /v3/{project_id}/volumes
	return pagination.NewPager(client, client.ServiceURL("volumes")+q.String(),
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
	var s []golangsdk.Link
	err := extract.IntoSlicePtr(r.BodyReader(), &s, "volumes_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s)
}

// ExtractVolumes extracts and returns Volumes. It is used while iterating over a volumes.List call.
func ExtractVolumes(r pagination.Page) ([]Volume, error) {
	var s []Volume
	err := ExtractVolumesInto(r, &s)
	return s, err
}

// ExtractVolumesInto similar to ExtractInto but operates on a `list` of volumes
func ExtractVolumesInto(r pagination.Page, v interface{}) error {
	return extract.IntoSlicePtr(r.(VolumePage).BodyReader(), v, "volumes")
}
