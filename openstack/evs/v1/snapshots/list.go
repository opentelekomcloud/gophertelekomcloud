package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"net/url"
)

type ListOpts struct {
	Name     string `q:"display_name"`
	Status   string `q:"status"`
	VolumeID string `q:"volume_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Snapshot, error) {
	var opts2 interface{} = opts
	query, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("snapshots")+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Snapshot
	err = extract.IntoSlicePtr(raw.Body, &res, "snapshots")
	return res, err
}
