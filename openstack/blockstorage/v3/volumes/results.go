package volumes

import (
	"encoding/json"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

type Attachment struct {
	AttachedAt   time.Time `json:"-"`
	AttachmentID string    `json:"attachment_id"`
	Device       string    `json:"device"`
	HostName     string    `json:"host_name"`
	ID           string    `json:"id"`
	ServerID     string    `json:"server_id"`
	VolumeID     string    `json:"volume_id"`
}

func (r *Attachment) UnmarshalJSON(b []byte) error {
	type tmp Attachment
	var s struct {
		tmp
		AttachedAt golangsdk.JSONRFC3339MilliNoZ `json:"attached_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Attachment(s.tmp)

	r.AttachedAt = time.Time(s.AttachedAt)

	return err
}

type Volume struct {
	// Unique identifier for the volume.
	ID string `json:"id"`
	// Current status of the volume.
	Status string `json:"status"`
	// Size of the volume in GB.
	Size int `json:"size"`
	// AvailabilityZone is which availability zone the volume is in.
	AvailabilityZone string `json:"availability_zone"`
	// The date when this volume was created.
	CreatedAt time.Time `json:"-"`
	// The date when this volume was last updated
	UpdatedAt time.Time `json:"-"`
	// Instances onto which the volume is attached.
	Attachments []Attachment `json:"attachments"`
	// Human-readable display name for the volume.
	Name string `json:"name"`
	// Human-readable description for the volume.
	Description string `json:"description"`
	// The type of volume to create, either SATA or SSD.
	VolumeType string `json:"volume_type"`
	// The ID of the snapshot from which the volume was created
	SnapshotID string `json:"snapshot_id"`
	// The ID of another block storage volume from which the current volume was created
	SourceVolID string `json:"source_volid"`
	// Arbitrary key-value pairs defined by the user.
	Metadata map[string]string `json:"metadata"`
	// UserID is the id of the user who created the volume.
	UserID string `json:"user_id"`
	// Indicates whether this is a bootable volume.
	Bootable string `json:"bootable"`
	// Encrypted denotes if the volume is encrypted.
	Encrypted bool `json:"encrypted"`
	// ReplicationStatus is the status of replication.
	ReplicationStatus string `json:"replication_status"`
	// ConsistencyGroupID is the consistency group ID.
	ConsistencyGroupID string `json:"consistencygroup_id"`
	// Multiattach denotes if the volume is multi-attach capable.
	Multiattach bool `json:"multiattach"`
}

func (r *Volume) UnmarshalJSON(b []byte) error {
	type tmp Volume
	var s struct {
		tmp
		CreatedAt golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
		UpdatedAt golangsdk.JSONRFC3339MilliNoZ `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Volume(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return err
}
