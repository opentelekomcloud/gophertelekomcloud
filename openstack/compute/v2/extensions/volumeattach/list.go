package volumeattach

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List returns a Pager that allows you to iterate over a collection of
// VolumeAttachments.
func List(client *golangsdk.ServiceClient, serverID string) ([]VolumeAttachment, error) {
	raw, err := client.Get(client.ServiceURL("servers", serverID, "os-volume_attachments"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []VolumeAttachment
	err = extract.IntoSlicePtr(raw.Body, &res, "volumeAttachments")
	return res, err
}
