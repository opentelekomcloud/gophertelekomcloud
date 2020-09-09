package ruletypes

import "github.com/opentelekomcloud/gophertelekomcloud"

func listRuleTypesURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("qos", "rule-types")
}
