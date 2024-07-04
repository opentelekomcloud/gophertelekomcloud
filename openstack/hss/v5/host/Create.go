package hss

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	// Server group name
	// Minimum: 1
	// Maximum: 128
	Name string `json:"group_name" required:"true"`
	// Server ID list
	// Array Length: 1 - 10000
	HostIds []string `json:"host_id_list" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v5/{project_id}/host-management/groups
	_, err = client.Post(client.ServiceURL("host-management", "groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"region": client.RegionID},
	})
	return err
}
