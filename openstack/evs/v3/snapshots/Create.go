package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// CreateOpts contains options for creating a Snapshot. This object is passed to
// the snapshots.Create function. For more information about these parameters,
// see the Snapshot object.
type CreateOpts struct {
	// Specifies the ID of the snapshot's source disk.
	VolumeID string `json:"volume_id" required:"true"`
	// Specifies the flag for forcibly creating a snapshot. The default value is false.
	// If this parameter is set to false and the disk is in the attaching state, the snapshot cannot be forcibly created.
	// If this parameter is set to true and the disk is in the attaching state, the snapshot can be forcibly created.
	Force bool `json:"force,omitempty"`
	// Specifies the snapshot name. The value can contain a maximum of 255 bytes.
	// NOTE
	// When creating a backup for a disk, a snapshot will be created and named with prefix autobk_snapshot_.
	// The EVS console has imposed operation restrictions on snapshots with prefix autobk_snapshot_. Therefore,
	// you are advised not to use autobk_snapshot_ as the name prefix for the snapshots you created.
	// Otherwise, the snapshots cannot be used normally.
	Name string `json:"name,omitempty"`
	// Specifies the snapshot description. The value can be null. The value can contain a maximum of 255 bytes.
	Description string `json:"description,omitempty"`
	// Specifies the snapshot metadata.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Create will create a new Snapshot based on the values in CreateOpts. To
// extract the Snapshot object from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Snapshot, error) {
	b, err := build.RequestBody(opts, "snapshot")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/snapshots
	raw, err := client.Post(client.ServiceURL("snapshots"), b, nil, nil)
	return extra(err, raw)
}
