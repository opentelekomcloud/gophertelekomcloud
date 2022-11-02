package volumes

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get retrieves the Volume with the provided ID.
func Get(client *golangsdk.ServiceClient, id string) (*Volume, error) {
	// GET /v3/{project_id}/volumes/{volume_id}
	raw, err := client.Get(client.ServiceURL("volumes", id), nil, nil)
	return extra(err, raw)
}
