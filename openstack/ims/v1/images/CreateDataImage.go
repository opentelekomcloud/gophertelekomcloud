package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v1/others"
)

type CreateDataImageOpts struct {
	// Specifies the image name.
	Name string `json:"name,omitempty"`
	// Provides supplementary information about the image. For detailed description, see Image Attributes. The value contains a maximum of 1024 characters and consists of only letters and digits. Carriage returns and angle brackets (< >) are not allowed. This parameter is left blank by default.
	Description string `json:"description,omitempty"`
	// Specifies the OS type.
	//
	// It can only be Windows or Linux. The default is Linux.
	OsType string `json:"os_type,omitempty"`
	// Specifies the URL of the external image file in the OBS bucket.
	//
	// The format is OBS bucket name:Image file name.
	//
	// NOTE:
	// The storage class of the OBS bucket must be Standard.
	ImageUrl string `json:"image_url,omitempty"`
	// Specifies the minimum size of the data disk.
	// Value range: 40 GB to 2048 GB
	MinDisk int `json:"min_disk" required:"true"`
	// Specifies the master key used for encrypting an image. For its value, see the Key Management Service User Guide.
	CmkId string `json:"cmk_id,omitempty"`
	// Specifies image tags. This parameter is left blank by default.
	//
	// For detailed parameter description, see Image Tag Data Formats.
	//
	// Use either tags or image_tags.
	Tags []string `json:"tags,omitempty"`
	// Lists the image tags. This parameter is left blank by default.
	//
	// Use either tags or image_tags.
	ImageTags []tags.ResourceTag `json:"image_tags,omitempty"`
}

// CreateDataImage This API is used to create a data disk image from a data disk image file uploaded to the OBS bucket. The API is an asynchronous one. If it is successfully called, the cloud service system receives the request. However, you need to use the asynchronous job query API to query the image creation status. For details, see Asynchronous Job Query.
func CreateDataImage(client *golangsdk.ServiceClient, opts CreateDataImageOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/cloudimages/dataimages/action
	raw, err := client.Post(client.ServiceURL("cloudimages", "dataimages", "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return others.ExtractJobId(err, raw)
}
