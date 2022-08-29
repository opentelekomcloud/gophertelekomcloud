package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// AllTenants will retrieve snapshots of all tenants/projects.
	AllTenants bool `q:"all_tenants"`
	// Name will filter by the specified snapshot name.
	Name string `q:"name"`
	// Status will filter by the specified status.
	Status string `q:"status"`
	// TenantID will filter by a specific tenant/project ID.
	// Setting AllTenants is required to use this.
	TenantID string `q:"project_id"`
	// VolumeID will filter by a specified volume ID.
	VolumeID string `q:"volume_id"`
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

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("snapshots")+query.String(), func(r pagination.PageResult) pagination.Page {
		return SnapshotPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

type SnapshotPage struct {
	pagination.LinkedPageBase
}

func (r SnapshotPage) IsEmpty() (bool, error) {
	volumes, err := ExtractSnapshots(r)
	return len(volumes) == 0, err
}

func (r SnapshotPage) NextPageURL() (string, error) {
	var s struct {
		Links []golangsdk.Link `json:"snapshots_links"`
	}

	err := extract.Into(r.Body, &s)
	if err != nil {
		return "", err
	}

	return golangsdk.ExtractNextURL(s.Links)
}

func ExtractSnapshots(r pagination.Page) ([]Snapshot, error) {
	var res struct {
		Snapshots []Snapshot `json:"snapshots"`
	}

	err := extract.Into(r.(SnapshotPage).Result.Body, &res)
	return res.Snapshots, err
}
