package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/common/tags"
)

type AddImageTagOpts struct {
	// Specifies the image ID.
	ImageId string `json:"-" required:"true"`
	// Specifies the tag to be added or updated.
	Tag tags.ResourceTag `json:"tag" required:"true"`
}

func AddImageTag(client *golangsdk.ServiceClient, opts AddImageTagOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v2/{project_id}/images/{image_id}/tags
	_, err = client.Post(client.ServiceURL("images", opts.ImageId, "tags"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
