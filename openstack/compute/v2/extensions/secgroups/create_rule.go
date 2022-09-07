package secgroups

import "github.com/opentelekomcloud/gophertelekomcloud"

// CreateRule will add a new rule to an existing security group (whose ID is
// specified in CreateRuleOpts). You have the option of controlling inbound
// traffic from either an IP range (CIDR) or from another security group.
func CreateRule(client *golangsdk.ServiceClient, opts CreateRuleOptsBuilder) (r CreateRuleResult) {
	b, err := opts.ToRuleCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(rootRuleURL(client), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
