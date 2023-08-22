package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePrivacyOpts struct {
	// URL to which the rule applies.
	Url string `json:"url" required:"true"`
	// Masked field
	// Enumeration values:
	// params
	// cookie
	// header
	// form
	Category string `json:"category" required:"true"`
	// Name of the masked field
	Name string `json:"index" required:"true"`
	// Rule description
	Description string `json:"description"`
}

// UpdatePrivacy is used to update the data masking rule list.
func UpdatePrivacy(client *golangsdk.ServiceClient, policyId, ruleId string, opts UpdatePrivacyOpts) (*PrivacyRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/privacy/{rule_id}
	raw, err := client.Put(client.ServiceURL("waf", "policy", policyId, "privacy", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res PrivacyRule
	return &res, extract.Into(raw.Body, &res)
}
