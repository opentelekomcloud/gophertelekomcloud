package compliance

import golangsdk "github.com/opentelekomcloud/gophertelekomcloud"

func DisableRule(client *golangsdk.ServiceClient, domainId, id string) (err error) {
	// POST /v1/resource-manager/domains/{domain_id}/policy-assignments/{policy_assignment_id}/disable
	_, err = client.Post(client.ServiceURL("resource-manager", "domains", domainId, "policy-assignments", id, "disable"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
