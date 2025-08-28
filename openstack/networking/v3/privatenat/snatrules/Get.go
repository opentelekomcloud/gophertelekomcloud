package snatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query details about a specified SNAT rule.
func Get(client *golangsdk.ServiceClient, ruleId string) (*SnatCommonResponse, error) {
	// GET /v3/{project_id}/private-nat/snat-rules/{snat_rule_id}
	raw, err := client.Get(client.ServiceURL("private-nat", "snat-rules", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res SnatCommonResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
