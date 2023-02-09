package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
)

type ImportImageQuickOpts struct {
	// Specifies the image name.
	Name string `json:"name" required:"true"`
	// Provides supplementary information about the image. For detailed description, see Image Attributes. The value contains a maximum of 1024 characters and consists of only letters and digits. Carriage returns and angle brackets (< >) are not allowed. This parameter is left blank by default.
	Description string `json:"description,omitempty"`
	// Specifies the OS version.
	// This parameter is valid if an external image file uploaded to the OBS bucket is used to create an image. For its value, see Values of Related Parameters.
	// When a data disk image created, the value can be Linux or Windows. The default is Linux.
	OsVersion string `json:"os_version,omitempty"`
	// Specifies the URL of the external image file in the OBS bucket.
	//
	// This parameter is mandatory if an external image file in the OBS bucket is used to create an image. The format is OBS bucket name:Image file name.
	//
	// NOTE:
	// The storage class of the OBS bucket must be Standard.
	ImageUrl string `json:"image_url" required:"true"`
	// Specifies the minimum size of the system disk in the unit of GB.
	//
	// This parameter is mandatory if an external image file in the OBS bucket is used to create an image. The value ranges from 1 to 1024.
	MinDisk int `json:"min_disk" required:"true"`
	// Lists the image tags. This parameter is left blank by default.
	//
	// Set either tags or image_tags.
	Tags []string `json:"tags,omitempty"`
	// Lists the image tags. The value is left blank by default.
	//
	// Set either tags or image_tags.
	ImageTags []tags.ResourceTag `json:"image_tags,omitempty"`
	// Specifies the image type.
	// The parameter value is ECS/BMS for system disk images. The default value is ECS.
	// The parameter value is DataImage for data disk images.
	Type string `json:"type,omitempty"`
}

func ImportImageQuick(client *golangsdk.ServiceClient, opts ImportImageQuickOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/cloudimages/quickimport/action
	raw, err2 := client.Post(client.ServiceURL("cloudimages", "quickimport", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return others.ExtractJobId(err2, raw)
}
