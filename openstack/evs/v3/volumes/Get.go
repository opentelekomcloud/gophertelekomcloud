package volumes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get retrieves the Volume with the provided ID.
func Get(client *golangsdk.ServiceClient, id string) (*Volume, error) {
	// GET /v3/{project_id}/volumes/{volume_id}
	raw, err := client.Get(client.ServiceURL("volumes", id), nil, nil)
	return extra(err, raw)
}

// Volume contains all the information associated with an OpenStack Volume.
type Volume struct {
	// Unique identifier for the volume.
	ID string `json:"id"`
	// Specifies the disk URI.
	Links []golangsdk.Link `json:"links"`
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
	// Specifies the disk type.
	// Currently, the value can be SSD, SAS, SATA, co-p1, uh-l1, GPSSD, or ESSD.
	// SSD: specifies the ultra-high I/O disk type.
	// SAS: specifies the high I/O disk type.
	// SATA: specifies the common I/O disk type.
	// co-p1: specifies the high I/O (performance-optimized I) disk type.
	// uh-l1: specifies the ultra-high I/O (latency-optimized) disk type.
	// GPSSD: specifies the general purpose SSD disk type.
	// ESSD: specifies the extreme SSD disk type.
	// Disks of the co-p1 and uh-l1 types are used exclusively for HPC ECSs and SAP HANA ECSs.
	VolumeType string `json:"volume_type"`
	// The ID of the snapshot from which the volume was created
	SnapshotID string `json:"snapshot_id"`
	// The ID of another block storage volume from which the current volume was created
	// Currently, this field is not supported by EVS.
	SourceVolID string `json:"source_volid"`
	// The backup ID, from which the volume was restored
	// This field is supported since 3.47 microversion
	BackupID string `json:"backup_id"`
	// Arbitrary key-value pairs defined by the user.
	Metadata map[string]string `json:"metadata"`
	// UserID is the id of the user who created the volume.
	UserID string `json:"user_id"`
	// Indicates whether this is a bootable volume.
	Bootable string `json:"bootable"`
	// Encrypted denotes if the volume is encrypted.
	// Currently, this field is not supported by EVS.
	Encrypted bool `json:"encrypted"`
	// ReplicationStatus is the status of replication.
	// Currently, this field is not supported by EVS.
	ReplicationStatus string `json:"replication_status"`
	// ConsistencyGroupID is the consistency group ID.
	ConsistencyGroupID string `json:"consistencygroup_id"`
	// Specifies whether the disk is shareable.
	Multiattach bool `json:"multiattach"`
	// Image metadata entries, only included for volumes that were created from an image, or from a snapshot of a volume originally created from an image.
	VolumeImageMetadata map[string]string `json:"volume_image_metadata"`
	// TenantID is the id of the project that owns the volume.
	TenantID string `json:"os-vol-tenant-attr:tenant_id"`
}

// UnmarshalJSON another unmarshalling function
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

func extra(err error, raw *http.Response) (*Volume, error) {
	if err != nil {
		return nil, err
	}

	var res Volume
	err = extract.IntoStructPtr(raw.Body, &res, "volume")
	return &res, err
}
