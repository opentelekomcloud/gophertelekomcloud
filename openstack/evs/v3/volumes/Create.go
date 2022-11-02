package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts contains options for creating a Volume. This object is passed to
// the volumes.Create function. For more information about these parameters, see the Volume object.
type CreateOpts struct {
	// Specifies the disk size, in GB. Its value can be as follows:
	// System disk: 1 GB to 1024 GB
	// Data disk: 10 GB to 32768 GB
	//
	// This parameter is mandatory when you create an empty disk.
	// You can specify the parameter value as required within the value range.
	//
	// This parameter is mandatory when you create the disk from a snapshot.
	// Ensure that the disk size is greater than or equal to the snapshot size.
	//
	// This parameter is mandatory when you create the disk from an image.
	// Ensure that the disk size is greater than or equal to the minimum disk capacity required by min_disk in the image attributes.
	Size int `json:"size,omitempty"`
	// Specifies the AZ where you want to create the disk. If the AZ does not exist, the disk will fail to create.
	AvailabilityZone string `json:"availability_zone"`
	// ConsistencyGroupID is the ID of a consistency group
	// Currently, this function is not supported.
	ConsistencyGroupID string `json:"consistencygroup_id,omitempty"`
	// Specifies the disk description. The value can contain a maximum of 255 bytes.
	Description string `json:"description,omitempty"`
	// Specifies the disk metadata. The length of the key or value in the metadata cannot exceed 255 bytes.
	// For details about metadata, see Parameters in the metadata field. The table lists some fields.
	// You can also specify other fields based on the disk creation requirements.
	// NOTE
	// Parameter values under metadata cannot be null.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Specifies the disk name. The value can contain a maximum of 255 bytes.
	Name string `json:"name,omitempty"`
	// Specifies the snapshot ID. If this parameter is specified, the disk is created from a snapshot.
	SnapshotID string `json:"snapshot_id,omitempty"`
	// Specifies the source disk ID. If this parameter is specified, the disk is cloned from an existing disk.
	// Currently, this function is not supported.
	SourceReplica string `json:"source_replica,omitempty"`
	// Specifies the source disk ID. If this parameter is specified, the disk is cloned from an existing disk.
	// Currently, this function is not supported.
	SourceVolID string `json:"source_volid,omitempty"`
	// Specifies the image ID. If this parameter is specified, the disk is created from an image.
	// NOTE BMS system disks cannot be created from BMS images.
	ImageID string `json:"imageRef,omitempty"`
	// Specifies the backup ID, from which you want to create the volume.
	// Create a volume from a backup is supported since 3.47 microversion
	BackupID string `json:"backup_id,omitempty"`
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
	// If the specified disk type is not available in the AZ, the disk will fail to create.
	// NOTE
	// If the disk is created from a snapshot, the volume_type field must be the same as that of the snapshot's source disk.
	VolumeType string `json:"volume_type,omitempty"`
	// Specifies whether the disk is shareable. The default value is false.
	// true: specifies a shared disk.
	// false: specifies a non-shared disk.
	Multiattach bool `json:"multiattach,omitempty"`
}

// Create will create a new Volume based on the values in CreateOpts.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Volume, error) {
	b, err := build.RequestBody(opts, "volume")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(createURL(client), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res Volume
	err = extract.IntoStructPtr(raw.Body, &res, "volume")
	return &res, err
}
