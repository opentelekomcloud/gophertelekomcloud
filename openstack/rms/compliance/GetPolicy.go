package compliance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetPolicy(client *golangsdk.ServiceClient, id string) (*PolicyDefinition, error) {
	// GET  /v1/resource-manager/policy-definitions/{policy_definition_id}
	raw, err := client.Get(client.ServiceURL(
		"resource-manager", "policy-definitions", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res PolicyDefinition
	err = extract.Into(raw.Body, &res)
	return &res, err
}
