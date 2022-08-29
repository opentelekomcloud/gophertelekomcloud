package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
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
}

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("snapshots")+query.String(), func(r pagination.PageResult) pagination.Page {
		return SnapshotPage{pagination.SinglePageBase(r)}
	})
}

type SnapshotPage struct {
	pagination.SinglePageBase
}

func (r SnapshotPage) IsEmpty() (bool, error) {
	volumes, err := ExtractSnapshots(r)
	return len(volumes) == 0, err
}

func ExtractSnapshots(r pagination.Page) ([]Snapshot, error) {
	var res struct {
		Snapshots []Snapshot `json:"snapshots"`
	}
	err := (r.(SnapshotPage)).ExtractInto(&res)
	return res.Snapshots, err
}
