package others

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CopyImageInRegionOpts struct {
	// Specifies the image ID.
	ImageId string `json:"-" required:"true"`
	// Specifies the image name.
	Name string `json:"name" required:"true"`
	// Specifies the encryption key. This parameter is left blank by default.
	CmkId string `json:"cmk_id,omitempty"`
	// Provides supplementary information about the image. For detailed description, see Image Attributes. The value contains a maximum of 1024 characters and consists of only letters and digits. Carriage returns and angle brackets (< >) are not allowed. This parameter is left blank by default.
	Description string `json:"description,omitempty"`
}

func CopyImageInRegion(client *golangsdk.ServiceClient, opts CopyImageInRegionOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/cloudimages/{image_id}/copy
	raw, err := client.Post(client.ServiceURL("cloudimages", opts.ImageId, "copy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return ExtractJobId(err, raw)
}
