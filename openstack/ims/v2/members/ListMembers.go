package members

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListMembers(client *golangsdk.ServiceClient, imageId string) (*ListMembersResponse, error) {
	// GET /v2/images/{image_id}/members
	raw, err := client.Get(client.ServiceURL("images", imageId, "members"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListMembersResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListMembersResponse struct {
	Members []Member `json:"members"`
	Schema  string   `json:"schema"`
}
