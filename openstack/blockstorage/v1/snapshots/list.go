package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Name     string `q:"display_name"`
	Status   string `q:"status"`
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
	err := extract.Into(r.(SnapshotPage), &res)
	return res.Snapshots, err
}
