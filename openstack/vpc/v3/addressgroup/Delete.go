package addressgroup

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// This function is used to delete an IP address group.
func Delete(client *golangsdk.ServiceClient, groupId string) error {
	// DELETE /v3/{project_id}/vpc/address-groups/{address_group_id}
	_, err := client.Delete(client.ServiceURL("address-groups", groupId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
