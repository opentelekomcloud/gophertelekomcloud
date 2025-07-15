package snatrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, id string) (*SnatRule, error) {
	// GET /v2.0/snat_rules/{snat_rule_id}
	raw, err := client.Get(client.ServiceURL("snat_rules", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res SnatRule
	return &res, extract.IntoStructPtr(raw.Body, &res, "snat_rule")
}
