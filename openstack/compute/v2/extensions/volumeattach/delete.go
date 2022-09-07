package volumeattach

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete requests the deletion of a previous stored VolumeAttachment from
// the server.
func Delete(client *golangsdk.ServiceClient, serverID, attachmentID string) (r DeleteResult) {
	raw, err := client.Delete(deleteURL(client, serverID, attachmentID), nil)
	return
}
