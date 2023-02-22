package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
)

type ExportImageOpts struct {
	// Specifies the image ID.
	ImageId string `json:"-" required:"true"`
	// Specifies the URL of the image file in the format of Bucket name:File name.
	//
	// Note
	//
	// The storage class of the OBS bucket must be Standard.
	BucketUrl string `json:"bucket_url" required:"true"`
	// Specifies the file format. The value can be qcow2, vhd, zvhd, or vmdk.
	FileFormat string `json:"file_format" required:"true"`
	// Whether to enable fast export. The value can be true or false.
	//
	// Note
	//
	// If fast export is enabled, file_format cannot be specified.
	IsQuickExport *bool `json:"is_quick_export,omitempty"`
}

func ExportImage(client *golangsdk.ServiceClient, opts ExportImageOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/cloudimages/{image_id}/file
	raw, err := client.Post(client.ServiceURL("cloudimages", opts.ImageId, "file"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return others.ExtractJobId(err, raw)
}
