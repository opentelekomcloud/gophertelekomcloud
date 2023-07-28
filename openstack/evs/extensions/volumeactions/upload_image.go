package volumeactions

import (
	"encoding/json"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UploadImageOpts struct {
	// Container format, may be bare, ofv, ova, etc.
	ContainerFormat string `json:"container_format,omitempty"`
	// Disk format, may be raw, qcow2, vhd, vdi, vmdk, etc.
	DiskFormat string `json:"disk_format,omitempty"`
	// The name of image that will be stored in glance.
	ImageName string `json:"image_name,omitempty"`
	// Force image creation, usable if volume attached to instance.
	Force bool `json:"force,omitempty"`
}

func UploadImage(client *golangsdk.ServiceClient, id string, opts UploadImageOpts) (*VolumeImage, error) {
	b, err := build.RequestBodyMap(opts, "os-volume_upload_image")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("volumes", id, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		VolumeImage VolumeImage `json:"os-volume_upload_image"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.VolumeImage, err
}

type VolumeImage struct {
	// The ID of a volume an image is created from.
	VolumeID string `json:"id"`
	// Container format, may be bare, ofv, ova, etc.
	ContainerFormat string `json:"container_format"`
	// Disk format, may be raw, qcow2, vhd, vdi, vmdk, etc.
	DiskFormat string `json:"disk_format"`
	// Human-readable description for the volume.
	Description string `json:"display_description"`
	// The ID of the created image.
	ImageID string `json:"image_id"`
	// Human-readable display name for the image.
	ImageName string `json:"image_name"`
	// Size of the volume in GB.
	Size int `json:"size"`
	// Current status of the volume.
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
