package volumeattach

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// VolumeAttachment contains attachment information between a volume and server.
type VolumeAttachment struct {
	// ID is a unique id of the attachment.
	ID string `json:"id"`
	// Device is what device the volume is attached as.
	Device string `json:"device"`
	// VolumeID is the ID of the attached volume.
	VolumeID string `json:"volumeId"`
	// ServerID is the ID of the instance that has the volume attached.
	ServerID string `json:"serverId"`
}

func extra(err error, raw *http.Response) (*VolumeAttachment, error) {
	if err != nil {
		return nil, err
	}

	var res VolumeAttachment
	err = extract.IntoStructPtr(raw.Body, &res, "volumeAttachment")
	return &res, err
}
