package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// Get retrieves a particular Configuration based on its unique ID.
func Get(client *golangsdk.ServiceClient, id string) (*Configuration, error) {
	// GET https://{Endpoint}/v3/{project_id}/configurations/{config_id}
	raw, err := client.Get(client.ServiceURL("configurations", id), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res Configuration
	err = extract.Into(raw.Body, &res)
	return &res, err
}
