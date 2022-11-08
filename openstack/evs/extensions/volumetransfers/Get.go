package volumetransfers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get retrieves the Transfer with the provided ID. To extract the Transfer object
// from the response, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (*Transfer, error) {
	// GET /v3/{project_id}/os-volume-transfer/{transfer_id}
	raw, err := client.Get(client.ServiceURL("os-volume-transfer", id), nil, nil)
	return extra(err, raw)
}
