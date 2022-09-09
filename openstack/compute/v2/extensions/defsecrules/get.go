package defsecrules

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get will return details for a particular default rule.
func Get(client *golangsdk.ServiceClient, id string) (*DefaultRule, error) {
	raw, err := client.Get(client.ServiceURL("os-security-group-default-rules", id), nil, nil)
	return extra(err, raw)
}
