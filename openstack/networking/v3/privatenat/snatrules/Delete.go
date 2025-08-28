package snatrules

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

// This function is used to delete a specified SNAT rule.
func Delete(client *golangsdk.ServiceClient, ruleId string) error {
	// DELETE /v3/{project_id}/private-nat/snat-rules/{snat_rule_id}
	_, err := client.Delete(client.ServiceURL("private-nat", "snat-rules", ruleId), &golangsdk.RequestOpts{
		OkCodes: []int{200, 204},
	})
	return err
}
