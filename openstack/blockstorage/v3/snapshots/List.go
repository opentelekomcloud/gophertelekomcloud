package snapshots

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	Marker       string `q:"marker"`
	Offset       int    `q:"offset"`
	Limit        int    `q:"limit"`
	Name         string `q:"name"`
	SortDir      string `q:"sort_dir"`
	Status       string `q:"status"`
	VolumeID     string `q:"volume_id"`
	NameLike     string `q:"name~"`
	StatusLike   string `q:"status~"`
	VolumeIDLike string `q:"volume_id~"`
	SortKey      string `q:"sort_key"`
	WithCount    bool   `q:"with_count"`
}

type ListResponse struct {
	Snapshots      []Snapshot       `json:"snapshots"`
	SnapshotsLinks []golangsdk.Link `json:"snapshots_links"`
	Count          int              `json:"count"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResponse, error) {
	q, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("snapshots") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SnapshotPage{NewPageResult: r}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}

	res, err := extractSnapshots(pages)
	if err != nil {
		return nil, err
	}
	if opts.WithCount {
		res.Count = len(res.Snapshots)
	}
	return res, nil
}

type SnapshotPage struct {
	pagination.NewPageResult
}

func (p SnapshotPage) NewNextPageURL() (string, error) {
	var res ListResponse
	if err := extract.Into(bytes.NewReader(p.Body), &res); err != nil {
		return "", err
	}
	return golangsdk.ExtractNextURL(res.SnapshotsLinks)
}

func (p SnapshotPage) NewIsEmpty() (bool, error) {
	res, err := extractSnapshots(p)
	if err != nil {
		return false, err
	}
	return len(res.Snapshots) == 0, nil
}

func extractSnapshots(page pagination.NewPage) (*ListResponse, error) {
	var res ListResponse
	if err := extract.Into(bytes.NewReader(page.(SnapshotPage).Body), &res); err != nil {
		return nil, err
	}
	if res.Snapshots == nil {
		res.Snapshots = []Snapshot{}
	}
	return &res, nil
}
