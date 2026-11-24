package resolverrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete is used to delete an endpoint rule.
// Send DELETE /v2.1/resolverrules/{resolverrule_id}
func Delete(client *golangsdk.ServiceClient, ruleID string) error {
	if ruleID == "" {
		return golangsdk.ErrMissingInput{Argument: "ruleID"}
	}

	_, err := client.Delete(v21URL(client, "resolverrules", ruleID), &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return err
}
