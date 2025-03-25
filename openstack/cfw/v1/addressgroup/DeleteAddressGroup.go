package addressgroup

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete an address group.
// groupId: Address group ID. It is the same as ID retuned while creating an address group.
func DeleteAddressGroup(client *golangsdk.ServiceClient, groupId string) error {
	// DELETE /v1/{project_id}/address-sets/{set_id}
	_, err := client.Delete(client.ServiceURL("address-sets", groupId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
