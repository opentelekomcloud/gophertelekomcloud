package dnatrules

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete a specified DNAT rule.
func Delete(client *golangsdk.ServiceClient, ruleId string) error {
	// DELETE /v3/{project_id}/private-nat/dnat-rules/{dnat_rule_id}
	_, err := client.Delete(client.ServiceURL("private-nat", "dnat-rules", ruleId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
