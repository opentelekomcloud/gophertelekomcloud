package defsecrules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List will return a collection of default rules.
func List(client *golangsdk.ServiceClient) ([]DefaultRule, error) {
	raw, err := client.Get(client.ServiceURL("os-security-group-default-rules"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []DefaultRule
	err = extract.IntoSlicePtr(raw.Body, &res, "security_group_default_rules")
	return res, err
}
