package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func UpdateDDos(client *golangsdk.ServiceClient, floatingIpId string, opts ConfigOpts) (*TaskResponse, error) {
	b, err := golangsdk.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/antiddos/{floating_ip_id}
	raw, err := client.Put(client.ServiceURL("antiddos", floatingIpId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	var res TaskResponse
	err = extract.Into(raw.Body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
