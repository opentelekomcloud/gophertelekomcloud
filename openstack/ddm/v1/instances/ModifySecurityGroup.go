package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ModifySecurityGroupOpts struct {
	// Security group ID. The default value is the original security group ID. You can change the value as required.
	SecurityGroupId string `json:"security_group_id" required:"true"`
}

// This function is used to modify the security group of a DDM instance.
func ModifySecurityGroup(client *golangsdk.ServiceClient, instanceId string, opts ModifySecurityGroupOpts) (*string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/instances/{instance_id}/modify-security-group
	raw, err := client.Put(client.ServiceURL("instances", instanceId, "modify-security-group"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ModifySecurityGroupResponse
	return &res.SecurityGroupId, extract.Into(raw.Body, &res)
}

type ModifySecurityGroupResponse struct {
	// Security group ID. The default value is the original security group ID. You can change the value as required.
	SecurityGroupId string `json:"security_group_id"`
}
