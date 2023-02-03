package volumeattach

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts specifies volume attachment creation or import parameters.
type CreateOpts struct {
	// Device is the device that the volume will attach to the instance as. Omit for "auto".
	Device string `json:"device,omitempty"`
	// VolumeID is the ID of the volume to attach to the instance.
	VolumeID string `json:"volumeId" required:"true"`
}

// Create requests the creation of a new volume attachment on the server.
func Create(client *golangsdk.ServiceClient, serverID string, opts CreateOpts) (*VolumeAttachment, error) {
	b, err := build.RequestBody(opts, "volumeAttachment")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("servers", serverID, "os-volume_attachments"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
