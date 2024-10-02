package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowDefaultConfig(client *golangsdk.ServiceClient) (*ConfigOpts, error) {
	// GET /v1/{project_id}/antiddos/default-config
	raw, err := client.Get(client.ServiceURL("antiddos", "default-config"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ConfigOpts
	err = extract.Into(raw.Body, &res)
	return &res, err
}
