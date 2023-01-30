package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type SetSecurityGroupOpts struct {
	InstanceId string `json:"-"`
	// Specifies the security group ID.
	SecurityGroupId string `json:"security_group_id" required:"true"`
}

func SetSecurityGroup(c *golangsdk.ServiceClient, opts SetSecurityGroupOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/instances/{instance_id}/security-group
	raw, err := c.Put(c.ServiceURL("instances", opts.InstanceId, "security-group"), b, nil,
		&golangsdk.RequestOpts{OkCodes: []int{200}})
	return extra(err, raw)
}
