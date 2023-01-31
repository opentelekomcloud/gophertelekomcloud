package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
)

type RegisterImageOpts struct {
	// Specifies the image ID.
	//
	// image_id is the ID of the image you created by invoking the API for creating image metadata. Registration may fail if you use other image IDs.
	//
	// After this API is invoked, you can check the image status with the image ID. When the image status changes to active, the image file is successfully registered. For details, see Querying Image Details (Native OpenStack API).
	ImageId string `json:"-" required:"true"`
	// Specifies the URL of the image file in the format of Bucket name:File name.
	//
	// Image files in the bucket can be in ZVHD, QCOW2, VHD, RAW, VHDX, QED, VDI, QCOW, ZVHD2, or VMDK format.
	//
	// Note
	//
	// The storage class of the OBS bucket must be Standard.
	ImageUrl string `json:"image_url" required:"true"`
}

func RegisterImage(client *golangsdk.ServiceClient, opts RegisterImageOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/cloudimages/{image_id}/upload
	raw, err := client.Put(client.ServiceURL("cloudimages", opts.ImageId, "upload"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return others.ExtractJobId(err, raw)
}
