package group

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// This function is used to delete a security group.
func Delete(client *golangsdk.ServiceClient, groupId string) error {
	// DELETE /v3/{project_id}/vpc/security-groups/{security_group_id}
	_, err := client.Delete(client.ServiceURL("security-groups", groupId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return err
}
