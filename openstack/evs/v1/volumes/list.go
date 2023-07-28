package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
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

func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Volume, error) {
	var opts2 interface{} = opts
	query, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("volumes")+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Volume
	err = extract.IntoSlicePtr(raw.Body, &res, "volumes")
	return res, err
}
