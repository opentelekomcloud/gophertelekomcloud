package compliance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

func EnableRule(client *golangsdk.ServiceClient, domainId, id string) (err error) {
	// POST /v1/resource-manager/domains/{domain_id}/policy-assignments/{policy_assignment_id}/enable
	_, err = client.Post(client.ServiceURL("resource-manager", "domains", domainId, "policy-assignments", id, "enable"), nil, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
