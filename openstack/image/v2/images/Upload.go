package images

import (
	"io"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Upload uploads an image file.
func Upload(client *golangsdk.ServiceClient, id string, data io.Reader) (err error) {
	// PUT /v2/images/{image_id}/file
	_, err = client.Put(client.ServiceURL("images", id, "file"), data, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/octet-stream"},
		OkCodes:     []int{204},
	})

	return
}
