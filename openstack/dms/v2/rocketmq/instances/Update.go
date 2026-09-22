package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Instance name. Consists of 4 to 64 characters including letters, digits,
	// and hyphens (-) and must start with a letter.
	Name string `json:"name,omitempty"`
	// Description of an instance. 0-1024 characters.
	Description *string `json:"description,omitempty"`
	// Security group ID.
	SecurityGroupID string `json:"security_group_id,omitempty"`
	// Whether to enable ACL.
	EnableACL *bool `json:"enable_acl,omitempty"`
	// Whether to enable public access.
	EnablePublicIP *bool `json:"enable_publicip,omitempty"`
	// IDs of the EIPs bound to the instance, separated by commas (,).
	// Mandatory if public access is enabled.
	PublicIpID string `json:"publicip_id,omitempty"`
	// Enterprise project.
	EnterpriseProjectID string `json:"enterprise_project_id,omitempty"`
}

// Update the information of a RocketMQ instance.
// Send PUT to /v2/{project_id}/instances/{instance_id}
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOpts) error {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", id).Build()
	if err != nil {
		return err
	}

	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Put(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
