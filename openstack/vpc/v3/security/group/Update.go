package group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Whether to only send the check request.
	// The value can be:
	// - true: The request will be checked and the security group will not be created.
	//         Check items include mandatory parameters, request format, and constraints.
	//         If the check fails, the system returns an error.
	//         If the check succeeds, response code 202 will be returned.
	// - false (default): A request will be sent and a security group will be created.
	DryRun bool `json:"dry_run,omitempty"`
	// Request body for creating a security group.
	SecurityGroup SecurityGroupUpdateOptions `json:"security_group" required:"true"`
}

type SecurityGroupUpdateOptions struct {
	// Security group name.
	// The value can contain 1 to 64 characters, including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name,omitempty"`
	// Description about the security group.
	// The value can contain up to 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description,omitempty"`
}

// This function is used to update a security group.
func Update(client *golangsdk.ServiceClient, groupId string, opts UpdateOpts) (*SecurityGroup, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/vpc/security-groups/{security_group_id}
	raw, err := client.Put(client.ServiceURL("security-groups", groupId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res SecurityGroupResponse
	return &res.SecurityGroup, extract.Into(raw.Body, &res)
}
