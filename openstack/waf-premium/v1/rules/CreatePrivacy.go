package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreatePrivacyOpts struct {
	// URL protected by the data masking rule.
	// The value must be in the standard URL format, for example, /admin.
	Url string `json:"url" required:"true"`
	// Masked field.
	// Enumeration values:
	// params
	// cookie
	// header
	// form
	Category string `json:"category" required:"true"`
	// Name of the masked field.
	Name string `json:"index" required:"true"`
	// Rule description.
	Description string `json:"description"`
}

// CreatePrivacy will create a data masking rule on the values in CreateOpts.
func CreatePrivacy(client *golangsdk.ServiceClient, policyId string, opts CreatePrivacyOpts) (*PrivacyRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/policy/{policy_id}/privacy
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "privacy"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res PrivacyRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type PrivacyRule struct {
	// Rule ID.
	ID string `json:"id"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// Time the rule is created. The value is a 13-digit timestamp in ms.
	CreatedAt int64 `json:"timestamp"`
	// Rule status. The value can be:
	// 0: The rule is disabled.
	// 1: The rule is enabled.
	Status *int `json:"status"`
	// URL protected by the data masking rule.
	Url string `json:"url"`
	// Masked field.
	// Enumeration values:
	// params
	// cookie
	// header
	// form
	Category string `json:"category"`
	// Name of the masked field.
	Name string `json:"index"`
	// Rule description.
	Description string `json:"description"`
}
