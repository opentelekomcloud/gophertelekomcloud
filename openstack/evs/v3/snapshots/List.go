package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToSnapshotListQuery() (string, error)
}

// ListOpts holds options for listing Snapshots. It is passed to the snapshots.List function.
type ListOpts struct {
	// AllTenants will retrieve snapshots of all tenants/projects.
	AllTenants bool `q:"all_tenants"`
	// Name will filter by the specified snapshot name.
	Name string `q:"name"`
	// Status will filter by the specified status.
	Status string `q:"status"`
	// VolumeID will filter by a specified volume ID.
	VolumeID string `q:"volume_id"`
	// Specifies the result sorting order. The default value is desc.
	// desc: indicates the descending order.
	// asc: indicates the ascending order.
	Sort string `q:"sort_dir"`
	// Specifies the sorting query by name (sort_key=name).
	// This parameter is supported when the request version is 3.30 or later. The default sorting order is the descending order.
	SortKey string `q:"sort_key"`
	// Requests a page size of items.
	Limit int `q:"limit"`
	// Used in conjunction with limit to return a slice of items.
	Offset int `q:"offset"`
	// Specifies the ID of the last record on the previous page. The returned value is the value of the item after this one.
	Marker string `q:"marker"`
	// Specifies to return parameter counts in the response. This parameter indicates the number of snapshots queried.
	// This parameter is in the with_count=true format and is supported when the request version is 3.45 or later.
	// This parameter can be used together with parameters marker, limit, sort_key, sort_dir, or offset in the table.
	// It cannot be used with other filter parameters.
	WithCount bool `q:"with_count"`
	// Specifies the fuzzy search by disk name. This parameter is supported when the request version is 3.31 or later.
	FuzzyName     string `q:"name~"`
	FuzzyStatus   string `q:"status~"`
	FuzzyVolumeID string `q:"volume_id~"`
}

// List returns Snapshots optionally limited by the conditions provided in ListOpts.
func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	// GET /v3/{project_id}/snapshots
	return pagination.NewPager(client, client.ServiceURL("snapshots")+q.String(), func(r pagination.PageResult) pagination.Page {
		return SnapshotPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// SnapshotPage is a pagination.Pager that is returned from a call to the List function.
type SnapshotPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a SnapshotPage contains no Snapshots.
func (r SnapshotPage) IsEmpty() (bool, error) {
	volumes, err := ExtractSnapshots(r)
	return len(volumes) == 0, err
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (r SnapshotPage) NextPageURL() (string, error) {
	var s []golangsdk.Link
	err := extract.IntoSlicePtr(r.BodyReader(), &s, "snapshots_links")
	if err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(s)
}

// ExtractSnapshots extracts and returns Snapshots. It is used while iterating over a snapshots.List call.
func ExtractSnapshots(r pagination.Page) ([]Snapshot, error) {
	var s []Snapshot
	err := extract.IntoSlicePtr(r.(SnapshotPage).BodyReader(), &s, "snapshots")
	return s, err
}
