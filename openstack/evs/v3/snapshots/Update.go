package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// UpdateOpts contain options for updating an existing Snapshot. This object is passed
// to the snapshots.Update function. For more information about the parameters, see
// the Snapshot object.
type UpdateOpts struct {
	// Specifies the snapshot name. The value can contain a maximum of 255 bytes.
	// NOTE
	// When creating a backup for a disk, a snapshot will be created and named with prefix autobk_snapshot_.
	// The EVS console has imposed operation restrictions on snapshots with prefix autobk_snapshot_.
	// Therefore, you are advised not to use autobk_snapshot_ as the name prefix for the snapshots you created.
	// Otherwise, the snapshots cannot be used normally.
	Name string `json:"name,omitempty"`
	// Specifies the snapshot description. The value can contain a maximum of 255 bytes.
	Description string `json:"description,omitempty"`
}

// Update will update the Snapshot with provided information. To extract the updated
// Snapshot from the response, call the Extract method on the UpdateResult.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*Snapshot, error) {
	b, err := build.RequestBody(opts, "snapshot")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/snapshots/{snapshot_id}
	raw, err := client.Put(client.ServiceURL("snapshots", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
