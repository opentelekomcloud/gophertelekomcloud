package dnatrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*DnatRule, error) {
	// GET /v2.0/dnat_rules/{dnat_rule_id}
	raw, err := client.Get(client.ServiceURL("dnat_rules", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res DnatRule
	return &res, extract.IntoStructPtr(raw.Body, &res, "dnat_rule")
}
