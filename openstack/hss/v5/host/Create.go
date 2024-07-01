package hss

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
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

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*GroupResp, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v5/{project_id}/host-management/groups
	raw, err := client.Post(client.ServiceURL("host-management", "groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res GroupResp

	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GroupResp struct {
	// Server group name
	Name string `json:"group_name"`
	// Server ID list
	HostIds []string `json:"host_id_list"`
}
