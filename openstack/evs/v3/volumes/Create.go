package volumes

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

type CreateOptsBuilder interface {
	ToVolumeCreateMap() (map[string]interface{}, error)
}

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
	Size int `json:"size" required:"true"`
	// Specifies the AZ where you want to create the disk. If the AZ does not exist, the disk will fail to create.
	AvailabilityZone string `json:"availability_zone" required:"true"`
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

// Metadata
// The preceding table provides only some parameters in metadata for your reference. You can also specify other fields based on the disk creation requirements.
// If the disk is created from a snapshot, __system__encrypted and __system__cmkid are not supported, and the newly created disk has the same encryption attribute as that of the snapshot's source disk.
// If the disk is created from an image, __system__encrypted and __system__cmkid are not supported, and the newly created disk has the same encryption attribute as that of the image.
// If the disk is created from a snapshot, hw:passthrough is not supported, and the newly created disk has the same device type as that of the snapshot's source disk.
// If the disk is created from an image, hw:passthrough is not supported, and the device type of newly created disk is VBD.
type Metadata struct {
	// Specifies the encryption field in metadata. The value can be 0 (not encrypted) or 1 (encrypted).
	// If this parameter does not exist, the disk will not be encrypted by default.
	SystemEncrypted string `json:"__system__encrypted"`
	// Specifies the encryption CMK ID in metadata. This parameter is used together with
	// __system__encrypted for encryption. The length of cmkid is fixed at 36 bytes.
	// NOTE
	// For details about how to obtain the CMK ID, see Querying the List of CMKs in the Key Management Service API Reference.
	SystemCmkId string `json:"__system__cmkid"`
	// If this parameter is set to true, the disk device type is SCSI, that is,
	// Small Computer System Interface (SCSI), which allows ECS OSs to directly
	// access the underlying storage media and supports SCSI reservation commands.
	// If this parameter is set to false, the disk device type will be VBD, which supports only simple SCSI read/write commands.
	// If this parameter does not appear, the disk device type is VBD.
	// NOTE
	// If parameter shareable is set to true and parameter hw:passthrough is not specified, shared VBD disks are created.
	Passthrough string `json:"hw:passthrough"`
	// If the disk is created from a snapshot and linked cloning needs to be used, set this parameter to 0.
	FullClone string `json:"full_clone"`
}

func (c CreateOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(c, "volume")
}

// Create will create a new Volume based on the values in CreateOpts.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (*Volume, error) {
	b, err := opts.ToVolumeCreateMap()
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/volumes
	raw, err := client.Post(client.ServiceURL("volumes"), b, nil, nil)
	return extra(err, raw)
}

// WaitForStatus will continually poll the resource, checking for a particular
// status. It will do this for the amount of seconds defined.
func WaitForStatus(c *golangsdk.ServiceClient, id, status string, secs int) error {
	return golangsdk.WaitFor(secs, func() (bool, error) {
		current, err := Get(c, id)
		if err != nil {
			return false, err
		}

		if current.Status == status {
			return true, nil
		}

		return false, nil
	})
}
