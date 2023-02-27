package images

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
)

type CreateOpts struct {
	// Specifies the image OS version. For the value range, see Values of Related Parameters.
	//
	// If this parameter is not specified, the value Other Linux(64 bit) will be used. In that case, the ECS creation using this image may fail, and the ECS created using this image may fail to run properly.
	OsVersion string `json:"__os_version"`
	// Specifies the container format.
	//
	// The default value is bare.
	ContainerFormat string `json:"container_format"`
	// Specifies the image format. The value can be zvhd2, vhd, zvhd, raw, or qcow2. The default value is zvhd2.
	DiskFormat string `json:"disk_format"`
	// Specifies the minimum disk space (GB) required for running the image. The value ranges from 1 GB to 1024 GB.
	//
	// The value of this parameter must be greater than the image system disk capacity. Otherwise, the ECS creation may fail.
	MinDisk int `json:"min_disk"`
	// Specifies the minimum memory size (MB) required for running the image. The parameter value depends on ECS specifications. The default value is 0.
	MinRam int `json:"min_ram"`
	// Specifies the image name. If this parameter is not specified, its value is empty by default. In that case, ECS creation using this image will fail. The name contains 1 to 255 characters. For detailed description, see Image Attributes. This parameter is left blank by default.
	Name string `json:"name"`
	// Lists the image tags. The tag contains 1 to 255 characters. The value is left blank by default.
	Tags []string `json:"tags"`
	// Specifies whether the image is available to other tenants.
	//
	// The default value is private. When creating image metadata, the value of visibility can be set to private only.
	Visibility string `json:"visibility"`
	// Specifies whether the image is protected. A protected image cannot be deleted. The default value is false.
	Protected bool `json:"protected"`
}

// Create implements create image request.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*images.ImageInfo, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/images
	raw, err := client.Post(client.ServiceURL("images"), b, nil, &golangsdk.RequestOpts{OkCodes: []int{201}})
	return extractImage(err, raw)
}
