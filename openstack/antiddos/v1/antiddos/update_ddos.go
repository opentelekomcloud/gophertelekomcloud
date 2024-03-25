package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func UpdateDDos(client *golangsdk.ServiceClient, floatingIpId string, opts ConfigOpts) (*TaskResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/antiddos/{floating_ip_id}
	raw, err := client.Put(client.ServiceURL("antiddos", floatingIpId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res TaskResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
