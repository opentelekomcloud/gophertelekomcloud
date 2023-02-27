package images

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
)

// Update implements image updated request.
func Update(client *golangsdk.ServiceClient, imageId string, opts []images.UpdateImageOpts) (*images.ImageInfo, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PATCH /v2/images/{image_id}
	raw, err := client.Patch(client.ServiceURL("images", imageId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/openstack-images-v2.1-json-patch"},
	})
	if err != nil {
		return nil, err
	}

	var res images.ImageInfo
	err = extract.Into(raw.Body, &res)
	return &res, err
}
