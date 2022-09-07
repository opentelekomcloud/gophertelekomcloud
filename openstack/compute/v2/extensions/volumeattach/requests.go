package volumeattach

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOptsBuilder allows extensions to add parameters to the Create request.
type CreateOptsBuilder interface {
	ToVolumeAttachmentCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies volume attachment creation or import parameters.
type CreateOpts struct {
	// Device is the device that the volume will attach to the instance as.
	// Omit for "auto".
	Device string `json:"device,omitempty"`

	// VolumeID is the ID of the volume to attach to the instance.
	VolumeID string `json:"volumeId" required:"true"`
}

// ToVolumeAttachmentCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToVolumeAttachmentCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "volumeAttachment")
}
