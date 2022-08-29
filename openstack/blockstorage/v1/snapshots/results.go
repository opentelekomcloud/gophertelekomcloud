package snapshots

import (
	"encoding/json"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type Snapshot struct {
	// Current status of the Snapshot.
	Status string `json:"status"`
	// Display name.
	Name string `json:"display_name"`
	// Instances onto which the Snapshot is attached.
	Attachments []string `json:"attachments"`
	// Logical group.
	AvailabilityZone string `json:"availability_zone"`
	// Is the Snapshot bootable?
	Bootable string `json:"bootable"`
	// Date created.
	CreatedAt time.Time `json:"-"`
	// Display description.
	Description string `json:"display_description"`
	// See VolumeType object for more information.
	VolumeType string `json:"volume_type"`
	// ID of the Snapshot from which this Snapshot was created.
	SnapshotID string `json:"snapshot_id"`
	// ID of the Volume from which this Snapshot was created.
	VolumeID string `json:"volume_id"`
	// User-defined key-value pairs.
	Metadata map[string]string `json:"metadata"`
	// Unique identifier.
	ID string `json:"id"`
	// Size of the Snapshot, in GB.
	Size int `json:"size"`
}

func (r *Snapshot) UnmarshalJSON(b []byte) error {
	type tmp Snapshot
	var s struct {
		tmp
		CreatedAt golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Snapshot(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)

	return err
}
