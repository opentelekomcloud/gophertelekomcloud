package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// admin-only option. Set it to true to see all tenant volumes.
	AllTenants bool `q:"all_tenants"`
	// List only volumes that contain Metadata.
	Metadata map[string]string `q:"metadata"`
	// List only volumes that have Name as the display name.
	Name string `q:"display_name"`
	// List only volumes that have a status of Status.
	Status string `q:"status"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) pagination.Pager {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return pagination.Pager{Err: err}
	}

	return pagination.NewPager(client, client.ServiceURL("volumes")+query.String(), func(r pagination.PageResult) pagination.Page {
		return VolumePage{pagination.SinglePageBase(r)}
	})
}

type VolumePage struct {
	pagination.SinglePageBase
}

func (r VolumePage) IsEmpty() (bool, error) {
	volumes, err := ExtractVolumes(r)
	return len(volumes) == 0, err
}

func ExtractVolumes(r pagination.Page) ([]Volume, error) {
	var res struct {
		Volumes []Volume `json:"volumes"`
	}
	err := (r.(VolumePage)).ExtractInto(&res)
	return res.Volumes, err
}
