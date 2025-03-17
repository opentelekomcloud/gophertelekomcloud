package blackwhitelist

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete a blacklist or whitelist rule.
// listId: Blacklist or whitelist ID. It is the same as ID retuned while creating a rule.
// We can also get it using GetBlacklistOrWhitelistRule.
func DeleteBlacklistOrWhitelistRule(client *golangsdk.ServiceClient, listId string) error {
	// DELETE /v1/{project_id}/black-white-list/{list_id}
	_, err := client.Delete(client.ServiceURL("black-white-list", listId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
