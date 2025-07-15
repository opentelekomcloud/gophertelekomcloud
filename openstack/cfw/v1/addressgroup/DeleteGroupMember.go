package addressgroup

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete a member from an address group.
// itemId: Address group member ID. It can be obtained using GetGroupMember function.
func DeleteGroupMember(client *golangsdk.ServiceClient, itemId string) error {
	// DELETE /v1/{project_id}/address-items/{item_id}
	_, err := client.Delete(client.ServiceURL("address-items", itemId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
