package cloudvolumes

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// ListOpts holds options for listing Volumes. It is passed to the volumes.List
// function.
type ListOpts struct {
	// Name will filter by the specified volume name.
	Name string `q:"name"`
	// Status will filter by the specified status.
	Status string `q:"status"`
	// Metadata will filter results based on specified metadata.
	Metadata string `q:"metadata"`
	// Specifies the disk ID.
	ID string `q:"id"`
	// Specifies the disk IDs. The parameter value is in the ids=["id1","id2",...,"idx"] format.
	// In the response, the ids value contains valid disk IDs only. Invalid disk IDs will be ignored.
	// Details about a maximum of 60 disks can be queried.
	// If parameters id and ids are both specified in the request, id will be ignored.
	IDs string `q:"ids"`
	// Specifies the AZ.
	AvailabilityZone string `q:"availability_zone"`
	// Specifies the ID of the DSS storage pool. All disks in the DSS storage pool can be filtered out.
	// Only precise match is supported.
	DedicatedStorageID string `q:"dedicated_storage_id"`
	// Specifies the name of the DSS storage pool. All disks in the DSS storage pool can be filtered out.
	// Fuzzy match is supported.
	DedicatedStorageName string `q:"dedicated_storage_name"`
	// Specifies whether the disk is shareable.
	//   true: specifies a shared disk.
	//   false: specifies a non-shared disk.
	Multiattach bool `q:"multiattach"`
	// Specifies the service type. Currently, the supported services are EVS, DSS, and DESS.
	ServiceType string `q:"service_type"`
	// Specifies the server ID.
	// This parameter is used to filter all the EVS disks that have been attached to this server.
	ServerID string `q:"server_id"`
	// Specifies the keyword based on which the returned results are sorted.
	// The value can be id, status, size, or created_at, and the default value is created_at.
	SortKey string `q:"sort_key"`
	// Specifies the result sorting order. The default value is desc.
	//   desc: indicates the descending order.
	//   asc: indicates the ascending order.
	SortDir string `q:"sort_dir"`
	// Specifies the disk type ID.
	// You can obtain the disk type ID in Querying EVS Disk Types.
	// That is, the id value in the volume_types parameter description table.
	VolumeTypeID string `q:"volume_type_id"`
	// Requests a page size of items.
	Limit int `q:"limit"`
	// Used in conjunction with limit to return a slice of items.
	Offset int `q:"offset"`
	// The ID of the last-seen item.
	Marker string `q:"marker"`
}

// List returns Volumes optionally limited by the conditions provided in ListOpts.
func List(client *golangsdk.ServiceClient, opts ListOpts) ([]Volume, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL("cloudvolumes", "detail") + q.String(),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SignatureKeyPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}

	var s []Volume

	err = extract.IntoSlicePtr(bytes.NewReader((pages.(SignatureKeyPage)).Body), &s, "volumes")
	return s, err
}

type SignatureKeyPage struct {
	pagination.NewSinglePageBase
}

// Attachment contains the disk attachment information
type Attachment struct {
	// Specifies the ID of the attachment information
	AttachmentID string `json:"attachment_id"`
	// Specifies the disk ID
	VolumeID string `json:"volume_id"`
	// Specifies the ID of the attached resource, equals to volume_id
	ResourceID string `json:"id"`
	// Specifies the ID of the server to which the disk is attached
	ServerID string `json:"server_id"`
	// Specifies the name of the host accommodating the server to which the disk is attached
	HostName string `json:"host_name"`
	// Specifies the device name
	Device string `json:"device"`
	// Specifies the time when the disk was attached. Time format: UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	AttachedAt string `json:"attached_at"`
}

// VolumeMetadata is an object that represents the metadata about the disk.
type VolumeMetadata struct {
	// Specifies the parameter that describes the encryption CMK ID in metadata.
	// This parameter is used together with __system__encrypted for encryption.
	// The length of cmkid is fixed at 36 bytes.
	SystemCmkID string `json:"__system__cmkid"`
	// Specifies the parameter that describes the encryption function in metadata. The value can be 0 or 1.
	//   0: indicates the disk is not encrypted.
	//   1: indicates the disk is encrypted.
	//   If this parameter does not appear, the disk is not encrypted by default.
	SystemEncrypted string `json:"__system__encrypted"`
	// Specifies the clone method. When the disk is created from a snapshot,
	// the parameter value is 0, indicating the linked cloning method.
	FullClone string `json:"full_clone"`
	// Specifies the parameter that describes the disk device type in metadata. The value can be true or false.
	//   If this parameter is set to true, the disk device type is SCSI, that is, Small Computer System
	//     Interface (SCSI), which allows ECS OSs to directly access the underlying storage media and supports SCSI
	//     reservation commands.
	//   If this parameter is set to false, the disk device type is VBD (the default type),
	//     that is, Virtual Block Device (VBD), which supports only simple SCSI read/write commands.
	//   If this parameter does not appear, the disk device type is VBD.
	HwPassthrough string `json:"hw:passthrough"`
	// Specifies the parameter that describes the disk billing mode in metadata.
	// If this parameter is specified, the disk is billed on a yearly/monthly basis.
	// If this parameter is not specified, the disk is billed on a pay-per-use basis.
	OrderID string `json:"orderID"`
	// Specifies the resource type about the disk.
	ResourceType string `json:"resourceType"`
	// Specifies the special code about the disk.
	ResourceSpecCode string `json:"resourceSpecCode"`
	// Specifies whether disk is read-only.
	ReadOnly string `json:"readonly"`
	// Specifies the attached mode about the disk.
	AttachedMode string `json:"attached_mode"`
}

// Link is an object that represents a link to which the disk belongs.
type Link struct {
	// Specifies the corresponding shortcut link.
	Href string `json:"href"`
	// Specifies the shortcut link marker name.
	Rel string `json:"rel"`
}

// Volume contains all the information associated with a Volume.
type Volume struct {
	// Unique identifier for the volume.
	ID string `json:"id"`
	// Human-readable display name for the volume.
	Name string `json:"name"`
	// Current status of the volume.
	Status string `json:"status"`
	// Size of the volume in GB.
	Size int `json:"size"`
	// Human-readable description for the volume.
	Description string `json:"description"`
	// The type of volume to create, either SATA or SSD.
	VolumeType string `json:"volume_type"`
	// AvailabilityZone is which availability zone the volume is in.
	AvailabilityZone string `json:"availability_zone"`
	// Instances onto which the volume is attached.
	Attachments []Attachment `json:"attachments"`
	// Specifies the disk URI.
	Links []Link `json:"links"`
	// The metadata of the disk image.
	ImageMetadata map[string]string `json:"volume_image_metadata"`
	// The ID of the snapshot from which the volume was created
	SnapshotID string `json:"snapshot_id"`
	// The ID of another block storage volume from which the current volume was created
	SourceVolID string `json:"source_volid"`
	// Specifies the ID of the tenant to which the disk belongs. The tenant ID is actually the project ID.
	OsVolTenantAttrTenantID string `json:"os-vol-tenant-attr:tenant_id"`
	// Specifies the service type. The value can be EVS, DSS or DESS.
	ServiceType string `json:"service_type"`
	// Indicates whether this is a bootable volume.
	Bootable string `json:"bootable"`
	// Multiattach denotes if the volume is multi-attach capable.
	Multiattach bool `json:"multiattach"`
	// Specifies the ID of the DSS storage pool accommodating the disk.
	DedicatedStorageID string `json:"dedicated_storage_id"`
	// Specifies the name of the DSS storage pool accommodating the disk.
	DedicatedStorageName string `json:"dedicated_storage_name"`
	// Encrypted denotes if the volume is encrypted.
	Encrypted bool `json:"encrypted"`
	// wwn of the volume.
	WWN string `json:"wwn"`
	// ReplicationStatus is the status of replication.
	ReplicationStatus string `json:"replication_status"`
	// ConsistencyGroupID is the consistency group ID.
	ConsistencyGroupID string `json:"consistencygroup_id"`
	// Arbitrary key-value pairs defined by the metadata field table.
	Metadata VolumeMetadata `json:"metadata"`
	// Arbitrary key-value pairs defined by the user.
	Tags map[string]string `json:"tags"`
	// UserID is the id of the user who created the volume.
	UserID string `json:"user_id"`
	// The date when this volume was created.
	CreatedAt string `json:"created_at"`
	// The date when this volume was last updated
	UpdatedAt string `json:"updated_at"`
}
