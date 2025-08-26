package dnatrules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to query details about a specified DNAT rule.
func Get(client *golangsdk.ServiceClient, ruleId string) (*DnatCommonResponse, error) {
	// GET /v3/{project_id}/private-nat/dnat-rules/{dnat_rule_id}
	raw, err := client.Get(client.ServiceURL("private-nat", "dnat-rules", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res DnatCommonResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
