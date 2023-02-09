package tags

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type DeleteImageTagOpts struct {
	// Specifies the image ID.
	ImageId string `json:"-" required:"true"`
	// Specifies the key of the tag to be deleted.
	Key string `json:"-" required:"true"`
}

func DeleteImageTag(client *golangsdk.ServiceClient, opts DeleteImageTagOpts) (err error) {
	// DELETE /v2/{project_id}/images/{image_id}/tags/{key}
	_, err = client.Delete(client.ServiceURL("images", opts.ImageId, "tags", opts.Key), openstack.StdRequestOpts())
	return
}
