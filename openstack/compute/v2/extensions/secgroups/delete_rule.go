package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// DeleteRule will permanently delete a rule from a security group.
func DeleteRule(client *golangsdk.ServiceClient, id string) (r DeleteRuleResult) {
	raw, err := client.Delete(resourceRuleURL(client, id), nil)
	return
}
