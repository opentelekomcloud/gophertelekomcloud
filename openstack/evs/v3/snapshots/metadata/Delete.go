package metadata

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, volumeId string, key string) (err error) {
	// DELETE /v3/{project_id}/volumes/{volume_id}/metadata/{key}
	_, err = client.Delete(client.ServiceURL("volumes", volumeId, "metadata", key), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
