package servicegroup

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete a member from an service group.
// itemId: Service group member ID. It can be obtained using GetGroupMember function.
func DeleteGroupMember(client *golangsdk.ServiceClient, itemId string) error {
	// DELETE /v1/{project_id}/service-items/{item_id}
	_, err := client.Delete(client.ServiceURL("service-items", itemId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
