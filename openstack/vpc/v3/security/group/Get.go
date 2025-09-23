package group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain details about a security group.
func Get(client *golangsdk.ServiceClient, groupId string) (*SecurityGroup, error) {
	// GET /v3/{project_id}/vpc/security-groups/{security_group_id}
	raw, err := client.Get(client.ServiceURL("security-groups", groupId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res SecurityGroupResponse
	err = extract.Into(raw.Body, &res)
	return &res.SecurityGroup, err
}
