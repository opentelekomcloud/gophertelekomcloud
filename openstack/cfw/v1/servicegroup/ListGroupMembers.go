package servicegroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query members in an service group.
// groupId: Service group ID. It is the same as ID retuned while creating an service group.
func ListGroupMembers(client *golangsdk.ServiceClient, groupId string) ([]GroupMemberRecord, error) {
	// GET /v1/{project_id}/service-items
	url, err := golangsdk.NewURLBuilder().WithEndpoints("service-items").WithQueryParams(&GetGroupMemberQueryParameters{
		SetID:  groupId,
		Limit:  1024,
		Offset: "0",
	}).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetGroupMembersResponse
	err = extract.Into(raw.Body, &res)
	return res.Data.Records, err
}
