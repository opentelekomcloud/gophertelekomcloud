package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type CreateOrUpdateTagsOpts struct {
	// Specifies the image ID.
	ImageId string `json:"image_id" required:"true"`
	// Specifies the tag.
	//
	// Use either tag or image_tag.
	Tag string `json:"tag,omitempty"`
	// Lists the image tags. For detailed description, see Image Tag Data Formats. This parameter is left blank by default.
	//
	// Use either tag or image_tag.
	ImageTag tags.ResourceTag `json:"image_tag,omitempty"`
}

func CreateOrUpdateTags(client *golangsdk.ServiceClient, opts CreateOrUpdateTagsOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT /v1/cloudimages/tags
	_, err = client.Put(client.ServiceURL("cloudimages", "tags"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
