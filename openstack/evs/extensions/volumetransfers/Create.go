package volumetransfers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts contains options for a Volume transfer.
type CreateOpts struct {
	// The ID of the volume to transfer.
	VolumeID string `json:"volume_id" required:"true"`
	// Specifies the disk transfer name. The value can contain a maximum of 255 bytes.
	Name string `json:"name,omitempty"`
}

// Create will create a volume tranfer request based on the values in CreateOpts.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Transfer, error) {
	b, err := build.RequestBody(opts, "transfer")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/os-volume-transfer
	raw, err := client.Post(client.ServiceURL("os-volume-transfer"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return extra(err, raw)
}
