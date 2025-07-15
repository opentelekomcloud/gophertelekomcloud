package compliance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetRule(client *golangsdk.ServiceClient, domainId, id string) (*PolicyRule, error) {
	// GET  /v1/resource-manager/domains/{domain_id}/policy-assignments/{policy_assignment_id}
	raw, err := client.Get(client.ServiceURL(
		"resource-manager", "domains", domainId, "policy-assignments", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res PolicyRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}
