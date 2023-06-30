package images

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteImageOpts struct {
	ImageId string `json:"-" required:"true"`
	// Specifies whether to delete the CSBS backups or CBR backups associated with a full-ECS image when the image is deleted. The value can be true or false.
	// true: When a full-ECS image is deleted, its CSBS backups or CBR backups are also deleted.
	// false: When a full-ECS image is deleted, its CSBS backups or CBR backups are not deleted.
	DeleteBackup bool `json:"delete_backup,omitempty"`
}

func Delete(client *golangsdk.ServiceClient, opts DeleteImageOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// DELETE /v2/images/{image_id}
	_, err = client.DeleteWithBody(client.ServiceURL("images", opts.ImageId), b, nil)
	return
}
