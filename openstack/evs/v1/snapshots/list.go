package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	Name     string `q:"display_name"`
	Status   string `q:"status"`
	VolumeID string `q:"volume_id"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Snapshot, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("snapshots").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Snapshot
	err = extract.IntoSlicePtr(raw.Body, &res, "snapshots")
	return res, err
}
