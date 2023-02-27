package images

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/images"
)

// Get implements image get request.
func Get(client *golangsdk.ServiceClient, id string) (*images.ImageInfo, error) {
	// GET /v2/images/{image_id}
	raw, err := client.Get(client.ServiceURL("images", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res images.ImageInfo
	err = extract.Into(raw.Body, &res)
	return &res, err
}
