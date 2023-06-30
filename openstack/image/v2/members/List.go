package members

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/ims/v2/members"
)

// List members returns list of members for specified image id.
func List(client *golangsdk.ServiceClient, imageId string) (*ListResponse, error) {
	// GET /v2/images/{image_id}/members
	raw, err := client.Get(client.ServiceURL("images", imageId, "members"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, res)
	return &res, err
}

type ListResponse struct {
	Members []members.Member `json:"members"`
	// Specifies the sharing schema.
	Schema string `json:"schema"`
}
