package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateImageFromDiskOpts Create a system disk image from a data disk.
type CreateImageFromDiskOpts struct {
	// Specifies the name of the system disk image.
	Name string `json:"name" required:"true"`
	// Specifies the data disk ID.
	VolumeId string `json:"volume_id" required:"true"`
	// Specifies the OS version.
	//
	// Set the parameter value based on Values of Related Parameters. Otherwise, the created system disk image may be unavailable.
	//
	// During the creation of a system disk image, if the OS can be detected from the data disk, the OS version in the data disk is used. In this case, the os_version value is invalid. If the OS can be detected from the data disk, the os_version value is used.
	OsVersion string `json:"os_version" required:"true"`
	// Specifies the image type.
	//
	// The value can be ECS, BMS, FusionCompute, or Ironic.
	//
	// ECS and FusionCompute: indicates an ECS image.
	// BMS and Ironic: indicates a BMS image.
	// The default value is ECS.
	Type string `json:"type,omitempty"`
	// Specifies the image description. This parameter is left blank by default. For details, see Image Attributes.
	//
	// The image description must meet the following requirements:
	//
	// Contains only letters and digits.
	// Cannot contain carriage returns and angle brackets (< >).
	// Cannot exceed 1024 characters.
	Description string `json:"description,omitempty"`
	// Specifies the minimum memory size (MB) required for running the image.
	//
	// The parameter value depends on the ECS specifications. The default value is 0.
	MinRam int `json:"min_ram,omitempty"`
	// Specifies the maximum memory size (MB) required for running the image.
	//
	// The parameter value depends on the ECS specifications. The default value is 0.
	MaxRam int `json:"max_ram,omitempty"`
	// Specifies tags of the image. This parameter is left blank by default.
	//
	// Use either tags or image_tags.
	Tags []string `json:"tags,omitempty"`
	// Specifies tags of the image. This parameter is left blank by default.
	//
	// Use either tags or image_tags.
	ImageTags []tags.ResourceTag `json:"image_tags,omitempty"`
}

// CreateImageFromDisk Constraints (Creating a System Disk Image Using a Data Disk)
// Before using a data disk to create a system disk image, ensure that an OS has been installed on the data disk and has been optimized. For details about the optimization, see "Optimizing a Windows Private Image" and "Optimizing a Linux Private Image" in the Image Management Service User Guide.
// The system cannot verify that an OS has been installed on the data disk. Therefore, ensure that the value of os_version is valid when creating a system disk image from the data disk. For details, see Values of Related Parameters.
func CreateImageFromDisk(client *golangsdk.ServiceClient, opts CreateImageFromDiskOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return cloudImages(client, err, b)
}
