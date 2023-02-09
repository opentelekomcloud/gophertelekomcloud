package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type ImportImageQuickFromFileOpts struct {
	// Specifies the image name.
	Name string `json:"name" required:"true"`
	// Provides supplementary information about the image.
	//
	// For detailed description, see Image Attributes.
	//
	// The value contains a maximum of 1024 characters and consists of only letters and digits. Carriage returns and angle brackets (< >) are not allowed. This parameter is left blank by default.
	Description string `json:"description,omitempty"`
	// Specifies the OS version.
	//
	// This parameter is valid if an external image file uploaded to the OBS bucket is used to create an image. For its value, see Values of Related Parameters.
	OsVersion string `json:"os_version" required:"true"`
	// Specifies the URL of the external image file in the OBS bucket.
	//
	// This parameter is mandatory if an external image file in the OBS bucket is used to create an image. The format is OBS bucket name:Image file name.
	//
	// NOTE:
	// The storage class of the OBS bucket must be Standard.
	ImageUrl string `json:"image_url" required:"true"`
	// Specifies the minimum size (GB) of the system disk.
	//
	// This parameter is mandatory if an external image file in the OBS bucket is used to create an image.
	// The value ranges from 1 to 1024 and must be greater than the size of the selected image file.
	MinDisk int `json:"min_disk" required:"true"`
	// Lists the image tags. This parameter is left blank by default.
	//
	// Set either tags or image_tags.
	Tags []string `json:"tags,omitempty"`
	// Lists the image tags. The value is left blank by default.
	//
	// Set either tags or image_tags.
	ImageTags []tags.ResourceTag `json:"image_tags,omitempty"`
	// Specifies the image type. The parameter value is ECS/BMS for system disk images. The default value is ECS.
	Type string `json:"type,omitempty"`
	// Specifies the image architecture type. Available values include:
	//
	// x86
	// arm
	// The default value is x86.
	//
	// NOTE:
	// If the image architecture is ARM, the boot mode is automatically changed to UEFI.
	Architecture string `json:"architecture,omitempty"`
}

// ImportImageQuickFromFile This API is used to quickly create a private image from an oversized external image file that has uploaded to the OBS bucket. Currently, only ZVHD2 and RAW image files are supported, and the size of an image file cannot exceed 1 TB.
//
// The fast image creation function is only available for image files in RAW or ZVHD2 format. For other formats of image files that are smaller than 128 GB, you are advised to import these files with the common method.
//
// The API is an asynchronous one. If it is successfully called, the cloud service system receives the request. However, you need to use the asynchronous job query API to query the image creation status. For details, see Asynchronous Job Query.
//
// Constraints
// Before importing image files, ensure that the file format is RAW or ZVHD2 and the following have been done:
// RAW image files have been optimized, and bitmap files have been generated.
// ZVHD2 image files have been optimized as required.
//
// For how to convert image file formats and generate a bitmap file, see section "Quickly Importing an Image File" in the Image Management Service User Guide.
func ImportImageQuickFromFile(client *golangsdk.ServiceClient, opts ImportImageQuickFromFileOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return quickImport(client, b)
}
