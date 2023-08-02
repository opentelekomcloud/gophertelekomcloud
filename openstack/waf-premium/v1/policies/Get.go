package policies

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*Policy, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Policy
	return &res, extract.Into(raw.Body, &res)
}
