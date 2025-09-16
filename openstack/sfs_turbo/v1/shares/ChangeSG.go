package shares

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ChangeSGOpts struct {
	// Specifies the change_security_group object.
	ChangeSecurityGroup SecurityGroupOpts `json:"change_security_group" required:"true"`
}

type SecurityGroupOpts struct {
	// Specifies the ID of the security group to be modified.
	SecurityGroupID string `json:"security_group_id" required:"true"`
}

// ChangeSG will change security group to a SFS Turbo based on the values in ChangeSGOpts.
func ChangeSG(client *golangsdk.ServiceClient, shareId string, opts ChangeSGOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST /v1/{project_id}/sfs-turbo/shares/{share_id}/action
	_, err = client.Post(client.ServiceURL("sfs-turbo", "shares", shareId, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return err
}
