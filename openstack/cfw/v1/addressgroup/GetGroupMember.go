package addressgroup

import (
	"fmt"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query a particular member in an address group.
// groupId: Address group ID. It is the same as ID retuned while creating an address group.
// ipAddress: // IP address of the group member.
func GetGroupMember(client *golangsdk.ServiceClient, groupId string, ipAddress string) (*GroupMemberRecord, error) {
	// GET /v1/{project_id}/address-items
	url, err := golangsdk.NewURLBuilder().WithEndpoints("address-items").WithQueryParams(&GetGroupMemberQueryParameters{
		SetID:   groupId,
		Limit:   1024,
		Offset:  "0",
		Address: ipAddress,
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
	if len(res.Data.Records) != 0 {
		return &res.Data.Records[0], nil
	}
	return nil, fmt.Errorf("no group member found")
}
