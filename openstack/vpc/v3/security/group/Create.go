package group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Whether to only send the check request.
	// The value can be:
	// - true: The request will be checked and the security group will not be created.
	//         Check items include mandatory parameters, request format, and constraints.
	//         If the check fails, the system returns an error.
	//         If the check succeeds, response code 202 will be returned.
	// - false (default): A request will be sent and a security group will be created.
	DryRun bool `json:"dry_run,omitempty"`
	// Request body for creating a security group.
	SecurityGroup SecurityGroupOptions `json:"security_group" required:"true"`
}

type SecurityGroupOptions struct {
	// Security group name.
	// The value can contain 1 to 64 characters, including letters, digits, underscores (_), hyphens (-), and periods (.).
	Name string `json:"name" required:"true"`
	// Description about the security group.
	// The value can contain up to 255 characters and cannot contain angle brackets (< or >).
	Description string `json:"description,omitempty"`
	// Enterprise project ID.
	// When creating a security group, associate an enterprise project ID with the security group.
	// The project ID can be 0 or a string that contains a maximum of 36 characters in UUID format with hyphens (-).
	// 0 indicates the default enterprise project.
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// Security group tags. Value range: 0 to 20 key-value pairs.
	Tags []TagOption `json:"tags,omitempty"`
}

type TagOption struct {
	// Tag key.
	// - Can contain up to 36 Unicode characters and cannot be left blank.
	// - Each key value of a resource must be unique.
	// - Value can contain: Letters, Digits, and special characters (_ @ -).
	Key string `json:"key" required:"true"`
	// Tag value.
	// - Can contain up to 43 Unicode characters and can be left blank.
	// - Value can contain: Letters, Digits, and special characters (_ @ -).
	Value string `json:"value" required:"true"`
}

// This function is used to create a security group.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*SecurityGroup, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v3/{project_id}/vpc/security-groups
	raw, err := client.Post(client.ServiceURL("security-groups"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200, 201},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res SecurityGroupResponse
	return &res.SecurityGroup, extract.Into(raw.Body, &res)
}
