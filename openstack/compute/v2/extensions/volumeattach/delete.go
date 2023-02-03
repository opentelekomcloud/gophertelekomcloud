package volumeattach

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete requests the deletion of a previous stored VolumeAttachment from the server.
func Delete(client *golangsdk.ServiceClient, serverID, attachmentID string) (err error) {
	_, err = client.Delete(client.ServiceURL("servers", serverID, "os-volume_attachments", attachmentID), nil)
	return
}
