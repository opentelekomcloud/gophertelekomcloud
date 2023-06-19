package volumeactions

import (
	"encoding/json"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UploadImageOpts struct {
	// Specifies the container type of the exported image.
	// The value can be ami, ari, aki, ovf, or bare. The default value is bare.
	ContainerFormat string `json:"container_format,omitempty"`
	// Specifies the format of the exported image.
	// The value can be vhd, zvhd, zvhd2, raw, or qcow2. The default value is zvhd2.
	DiskFormat string `json:"disk_format,omitempty"`
	// Specifies the name of the exported image.
	// The name cannot start or end with space.
	// The name contains 1 to 128 characters.
	// The name contains the following characters: uppercase letters, lowercase letters, digits,
	// and special characters, such as hyphens (-), periods (.), underscores (_), and spaces.
	ImageName string `json:"image_name,omitempty"`
	// Specifies whether to forcibly export the image. The default value is false.
	// If force is set to false and the disk is in the in-use state, the image cannot be forcibly exported.
	// If force is set to true and the disk is in the in-use state, the image can be forcibly exported.
	Force bool `json:"force,omitempty"`
	// Specifies the OS type of the exported image. Currently, only windows and linux are supported. The default value is linux.
	// There are two underscores (_) in front of os and one underscore (_) after os.
	// This parameter setting takes effect only when the __os_type field is not included in volume_image_metadata and the disk status is available.
	// If this parameter is not specified, default value linux is used as the OS type of the image.
	OSType string `json:"__os_type,omitempty"`
}

func UploadImage(client *golangsdk.ServiceClient, id string, opts UploadImageOpts) (*VolumeImage, error) {
	b, err := build.RequestBody(opts, "os-volume_upload_image")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res VolumeImage
	err = extract.IntoStructPtr(raw.Body, &res, "os-volume_upload_image")
	return &res, err
}

type VolumeImage struct {
	// The ID of a volume an image is created from.
	VolumeID string `json:"id"`
	// Specifies the container type of the exported image.
	// The value can be ami, ari, aki, ovf, or bare. The default value is bare.
	ContainerFormat string `json:"container_format"`
	// Specifies the format of the exported image.
	// The value can be vhd, zvhd, zvhd2, raw, or qcow2. The default value is vhd.
	DiskFormat string `json:"disk_format"`
	// Human-readable description for the volume.
	Description string `json:"display_description"`
	// Specifies the ID of the exported image.
	ImageID string `json:"image_id"`
	// Specifies the name of the exported image.
	ImageName string `json:"image_name"`
	// Size of the volume in GB.
	Size int `json:"size"`
	// Specifies the disk status after the image is exported. The correct value is uploading.
	Status string `json:"status"`
	// The date when this volume was last updated.
	UpdatedAt time.Time `json:"-"`
	// Volume type object of used volume.
	VolumeType ImageVolumeType `json:"volume_type"`
}

func (r *VolumeImage) UnmarshalJSON(b []byte) error {
	type tmp VolumeImage
	var s struct {
		tmp
		UpdatedAt golangsdk.JSONRFC3339MilliNoZ `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = VolumeImage(s.tmp)

	r.UpdatedAt = time.Time(s.UpdatedAt)
	return err
}

type ImageVolumeType struct {
	// The ID of a volume type.
	ID string `json:"id"`
	// Human-readable display name for the volume type.
	Name string `json:"name"`
	// Human-readable description for the volume type.
	Description string `json:"display_description"`
	// Flag for public access.
	IsPublic bool `json:"is_public"`
	// Extra specifications for volume type.
	// volumetypes.ExtraSpecs
	ExtraSpecs map[string]interface{} `json:"extra_specs"`
	// ID of quality of service specs.
	QosSpecsID string `json:"qos_specs_id"`
	// Flag for deletion status of volume type.
	Deleted bool `json:"deleted"`
	// The date when volume type was deleted.
	DeletedAt time.Time `json:"-"`
	// The date when volume type was created.
	CreatedAt time.Time `json:"-"`
	// The date when this volume was last updated.
	UpdatedAt time.Time `json:"-"`
}

func (r *ImageVolumeType) UnmarshalJSON(b []byte) error {
	type tmp ImageVolumeType
	var s struct {
		tmp
		CreatedAt golangsdk.JSONRFC3339MilliNoZ `json:"created_at"`
		UpdatedAt golangsdk.JSONRFC3339MilliNoZ `json:"updated_at"`
		DeletedAt golangsdk.JSONRFC3339MilliNoZ `json:"deleted_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = ImageVolumeType(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)
	r.DeletedAt = time.Time(s.DeletedAt)

	return err
}
