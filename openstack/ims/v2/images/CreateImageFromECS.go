package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
)

type CreateImageFromECSOpts struct {
	// Specifies the name of the system disk image. For detailed description, see Image Attributes.
	Name string `json:"name" required:"true"`
	// Specifies the image description. For detailed description, see Image Attributes. The value contains a maximum of 1024 characters and consists of only letters and digits. Carriage returns and angle brackets (< >) are not allowed. This parameter is left blank by default.
	Description string `json:"description,omitempty"`
	// Specifies the ID of the ECS used to create the image.
	//
	// To obtain the ECS ID, perform the following operations:
	//
	// Log in to management console.
	// Under Computing, click Elastic Cloud Server.
	// In the ECS list, click the name of the ECS and view its ID.
	InstanceId string `json:"instance_id,omitempty"`
	// Specifies the data disk information to be converted. This parameter is mandatory when the data disk of an ECS is used to create a private data disk image. For details, see Table 1.
	//
	// If the ECS data disk is not used to create a data disk image, the parameter is empty by default.
	//
	// NOTE:
	// When you create a data disk image using a data disk, if other parameters (such as name, description, and tags) in this table have values, the system uses the value of data_images. You cannot specify instance_id.
	DataImages []ECSDataImage `json:"data_images,omitempty"`
	// Specifies tags of the image. This parameter is left blank by default.
	//
	// Use either tags or image_tags.
	Tags []string `json:"tags,omitempty"`
	// Specifies tags of the image. This parameter is left blank by default.
	//
	// Use either tags or image_tags.
	ImageTags []tags.ResourceTag `json:"image_tags,omitempty"`
	// Specifies the maximum memory of the image in the unit of MB.
	MaxRam int `json:"max_ram,omitempty"`
	// Specifies the minimum memory of the image in the unit of MB. The default value is 0, indicating that the memory is not restricted.
	MinRam int `json:"min_ram,omitempty"`
}

type ECSDataImage struct {
	// Specifies the name of a data disk image.
	Name string `json:"name" required:"true"`
	// Specifies the data disk ID.
	VolumeId string `json:"volume_id" required:"true"`
	// Specifies the data disk description.
	Description string `json:"description,omitempty"`
	// Specifies the data disk image tag.
	Tags []string `json:"tags,omitempty"`
}

// CreateImageFromECS This API is used to create a private image. The following methods are supported:
//
// Create a system or data disk image from an ECS.
// Create a system disk image from an external image file uploaded to an OBS bucket.
// Create a system disk image from a data disk.
// The API is an asynchronous one. If it is successfully called, the cloud service system receives the request. However, you need to use the asynchronous job query API to query the image creation status. For details, see Asynchronous Job Query.
//
// You cannot export public images (such as Windows, SUSE Linux, Red Hat Linux, Oracle Linux, and Ubuntu) or private images created using these public images.
func CreateImageFromECS(client *golangsdk.ServiceClient, opts CreateImageFromECSOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return cloudImages(client, b)
}

func cloudImages(client *golangsdk.ServiceClient, b *build.Body) (*string, error) {
	// POST /v2/cloudimages/action
	raw, err := client.Post(client.ServiceURL("cloudimages", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return others.ExtractJobId(err, raw)
}
