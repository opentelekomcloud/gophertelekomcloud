package volumeattach

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns public data about a previously created VolumeAttachment.
func Get(client *golangsdk.ServiceClient, serverID, attachmentID string) (r GetResult) {
	raw, err := client.Get(client.ServiceURL("servers", serverID, "os-volume_attachments", attachmentID), nil, nil)
	return
}
