package servicegroup

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query a particular member in an service group.
// groupId: Service group ID. It is the same as ID retuned while creating an service group.
// ipService: // IP service of the group member.
func GetGroupMember(client *golangsdk.ServiceClient, groupId, sourcePort, destPort string, protocol int) (*GroupMemberRecord, error) {
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
	if err != nil {
		return nil, err
	}
	for _, groupMember := range res.Data.Records {
		if groupMember.Protocol == protocol && groupMember.SourcePort == sourcePort && groupMember.DestPort == destPort {
			return &groupMember, nil
		}
	}
	return nil, fmt.Errorf("no group member found")
}
