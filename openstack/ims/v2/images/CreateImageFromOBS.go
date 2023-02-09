package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

// CreateImageFromOBSOpts Create a system disk image from an external image file uploaded to an OBS bucket.
type CreateImageFromOBSOpts struct {
	// Specifies the name of the system disk image. For detailed description, see Image Attributes.
	Name string `json:"name" required:"true"`
	// Specifies the image description. For detailed description, see Image Attributes. The value contains a maximum of 1024 characters and consists of only letters and digits. Carriage returns and angle brackets (< >) are not allowed. This parameter is left blank by default.
	Description string `json:"description,omitempty"`
	// Specifies the OS type. The value can be Linux, Windows, or Other.
	OsType string `json:"os_type,omitempty"`
	// Specifies the OS version.
	//
	// This parameter is valid if an external image file uploaded to the OBS bucket is used to create an image. For its value, see Values of Related Parameters.
	//
	// NOTE:
	// This parameter is mandatory when the value of is_quick_import is true, that is, a system disk image is imported using the quick import method.
	OsVersion string `json:"os_version,omitempty"`
	// Specifies the URL of the external image file in the OBS bucket.
	//
	// This parameter is mandatory if an external image file in the OBS bucket is used to create an image. The format is OBS bucket name:Image file name.
	//
	// To obtain an OBS bucket name:
	// Log in to the management console and choose Storage > Object Storage Service.
	// All OBS buckets are displayed in the list.
	//
	// Filter the OBS buckets by region and locate the target bucket in the current region.
	// To obtain an OBS image file name:
	// Log in to the management console and choose Storage > Object Storage Service.
	// All OBS buckets are displayed in the list.
	//
	// Filter the OBS buckets by region and locate the target bucket in the current region.
	// Click the name of the target bucket to go to the bucket details page.
	// In the navigation pane on the left, choose Objects to display objects in the OBS bucket and then locate the external image file used to create an image.
	// NOTE:
	// The storage class of the OBS bucket must be Standard.
	ImageUrl string `json:"image_url" required:"true"`
	// Specifies the minimum size of the system disk in the unit of GB.
	//
	// This parameter is mandatory if an external image file in the OBS bucket is used to create an image. The value ranges from 1 GB to 1024 GB.
	MinDisk int `json:"min_disk" required:"true"`
	// Specifies whether automatic configuration is enabled.
	//
	// The value can be true or false.
	//
	// If automatic configuration is required, set the value to true. Otherwise, set the value to false The default value is false.
	//
	// For details about automatic configuration, see Creating a Linux System Disk Image from an External Image File > Registering an External Image File as a Private Image (Linux) in Image Management Service User Guide.
	IsConfig bool `json:"is_config,omitempty"`
	// Specifies the master key used for encrypting an image. For its value, see the Key Management Service User Guide.
	CmkId string `json:"cmk_id,omitempty"`
	// Specifies tags of the image. This parameter is left blank by default.
	// Use either tags or image_tags.
	Tags []string `json:"tags,omitempty"`
	// Specifies tags of the image. This parameter is left blank by default.
	// Use either tags or image_tags.
	ImageTags []tags.ResourceTag `json:"image_tags,omitempty"`
	// Specifies the maximum memory of the image in the unit of MB.
	MaxRam int `json:"max_ram,omitempty"`
	// Specifies the minimum memory of the image in the unit of MB. The default value is 0, indicating that the memory is not restricted.
	MinRam int `json:"min_ram,omitempty"`
	// Specifies the data disk information to be imported.
	//
	// An external image file can contain a maximum of three data disks. In this case, one system disk and three data disks will be created.
	//
	// For details, see Table 2.
	//
	// NOTE:
	// If a data disk image file is used to create a data disk image, the OS type of the data disk image must be the same as that of the system disk image.
	// If other parameters (such as name, description, and tags) in Table 2 are set, the system uses the values in data_images.
	DataImages []OBSDataImage `json:"data_images,omitempty"`
	// Specifies whether to use the quick import method to import a system disk image.
	// For details about the restrictions on quick import of image files, see Importing an Image File Quickly.
	IsQuickImport bool `json:"is_quick_import,omitempty"`
}

type OBSDataImage struct {
	// Specifies the image name.
	Name string `json:"name,omitempty"`
	// Specifies the enterprise project that the image belongs to. The value is left blank by default.
	//
	// The value contains a maximum of 1024 characters and consists of only letters and digits. Carriage returns and angle brackets (< >) are not allowed.
	Description string `json:"description,omitempty"`
	// Specifies the URL of the external image file in the OBS bucket.
	//
	// The format is OBS bucket name:Image file name.
	//
	// To obtain an OBS bucket name:
	// Log in to the management console and choose Storage > Object Storage Service.
	// All OBS buckets are displayed in the list.
	//
	// Filter the OBS buckets by region and locate the target bucket in the current region.
	// To obtain an OBS image file name:
	// Log in to the management console and choose Storage > Object Storage Service.
	// All OBS buckets are displayed in the list.
	//
	// Filter the OBS buckets by region and locate the target bucket in the current region.
	// Click the name of the target bucket to go to the bucket details page.
	// In the navigation pane on the left, choose Objects to display objects in the OBS bucket and then locate the external image file used to create an image.
	// NOTE:
	// The storage class of the OBS bucket must be Standard.
	ImageUrl string `json:"image_url" required:"true"`
	// Specifies the data disk ID.
	VolumeId string `json:"volume_id"`
	// Specifies the minimum size of the data disk.
	// Unit: GB
	// Value range: 1–2048
	MinDisk int `json:"min_disk" required:"true"`
	// Specifies whether an image file is imported quickly to create a data disk image.
	IsQuickImport bool `json:"is_quick_import,omitempty"`
	// Specifies tags of the image. This parameter is left blank by default.
	// Use either tags or image_tags.
	Tags []string `json:"tags,omitempty"`
	// Specifies tags of the image. This parameter is left blank by default.
	// Use either tags or image_tags.
	ImageTags []tags.ResourceTag `json:"image_tags,omitempty"`
}

func CreateImageFromOBS(client *golangsdk.ServiceClient, opts CreateImageFromOBSOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return cloudImages(client, err, b)
}
