package servicegroup

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete an service group.
// groupId: Service group ID. It is the same as ID retuned while creating an service group.
func DeleteServiceGroup(client *golangsdk.ServiceClient, groupId string) error {
	// DELETE /v1/{project_id}/service-sets/{set_id}
	_, err := client.Delete(client.ServiceURL("service-sets", groupId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
