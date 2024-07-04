package hss

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Server group name
	// Minimum: 1
	// Maximum: 128
	Name string `json:"group_name,omitempty"`
	// Server group ID
	ID string `json:"group_id" required:"true"`
	// Server ID list
	// Array Length: 1 - 10000
	HostIds []string `json:"host_id_list,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*GroupResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v5/{project_id}/host-management/groups
	raw, err := client.Put(client.ServiceURL("host-management", "groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"region": client.RegionID},
	})
	if err != nil {
		return nil, err
	}

	var res GroupResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}
