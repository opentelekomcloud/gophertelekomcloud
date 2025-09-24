package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// This function is used to delete a security group rule.
func Delete(client *golangsdk.ServiceClient, ruleId string) error {
	// DELETE /v3/{project_id}/vpc/security-group-rules/{security_group_rule_id}
	_, err := client.Delete(client.ServiceURL("security-group-rules", ruleId), nil)
	return err
}
