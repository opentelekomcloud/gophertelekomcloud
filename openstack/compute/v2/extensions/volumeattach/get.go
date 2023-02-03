package volumeattach

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get returns public data about a previously created VolumeAttachment.
func Get(client *golangsdk.ServiceClient, serverID, attachmentID string) (*VolumeAttachment, error) {
	raw, err := client.Get(client.ServiceURL("servers", serverID, "os-volume_attachments", attachmentID), nil, nil)
	return extra(err, raw)
}
