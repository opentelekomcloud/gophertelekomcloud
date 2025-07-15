package acl

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete an ACL rule.
func DeleteACLRule(client *golangsdk.ServiceClient, ruleId string) error {
	// DELETE /v1/{project_id}/acl-rule/{acl_rule_id}
	_, err := client.Delete(client.ServiceURL("acl-rule", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return err
}
