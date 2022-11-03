package metadata

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Get(client *golangsdk.ServiceClient, volumeId string) (map[string]string, error) {
	// GET /v3/{project_id}/volumes/{volume_id}/metadata
	raw, err := client.Get(client.ServiceURL("volumes", volumeId, "metadata"), nil, nil)
	return extraMetadata(err, raw)
}
