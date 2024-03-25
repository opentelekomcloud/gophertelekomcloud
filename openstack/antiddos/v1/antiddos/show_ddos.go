package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowDDos(client *golangsdk.ServiceClient, floatingIpId string) (*ConfigOpts, error) {
	// GET /v1/{project_id}/antiddos/{floating_ip_id}
	raw, err := client.Get(client.ServiceURL("antiddos", floatingIpId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ConfigOpts
	err = extract.Into(raw.Body, &res)
	return &res, err
}
