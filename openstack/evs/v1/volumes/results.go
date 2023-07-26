package volumes

import (
	"encoding/json"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type Volume struct {
	// Current status of the volume.
	Status string `json:"status"`
	// Human-readable display name for the volume.
	Name string `json:"display_name"`
	// Instances onto which the volume is attached.
	Attachments []map[string]any `json:"attachments"`
	// This parameter is no longer used.
	AvailabilityZone string `json:"availability_zone"`
	// Indicates whether this is a bootable volume.
	Bootable string `json:"bootable"`
	// The date when this volume was created.
	CreatedAt time.Time `json:"-"`
	// Human-readable description for the volume.
	Description string `json:"display_description"`
	// The type of volume to create, either SATA or SSD.
	VolumeType string `json:"volume_type"`
	// The ID of the snapshot from which the volume was created
	SnapshotID string `json:"snapshot_id"`
	// The ID of another block storage volume from which the current volume was created
	SourceVolID string `json:"source_volid"`
	// Arbitrary key-value pairs defined by the user.
	Metadata map[string]string `json:"metadata"`
	// Unique identifier for the volume.
	ID string `json:"id"`
	// Size of the volume in GB.
	Size int `json:"size"`
}

func (r *Volume) UnmarshalJSON(b []byte) error {
	type tmp Volume
	var s struct {
		tmp
		CreatedAt golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Volume(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)

	return err
}
