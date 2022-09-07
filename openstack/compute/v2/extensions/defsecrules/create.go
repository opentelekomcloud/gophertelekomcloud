package defsecrules

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create is the operation responsible for creating a new default rule.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRuleCreateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("os-security-group-default-rules"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
