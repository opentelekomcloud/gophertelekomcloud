package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateKnownAttackSourceOpts struct {
	// Type of the know attack source rule.
	// Enumeration values:
	// long_ip_block
	// long_cookie_block
	// long_params_block
	// short_ip_block
	// short_cookie_block
	// short_params_block
	Category string `json:"category" required:"true"`
	// Block duration, in seconds. If prefix long is selected for the rule type,
	// the value for block_time ranges from 301 to 1800.
	// If prefix short is selected for the rule type,
	// the value for block_time ranges from 0 to 300.
	BlockTime int `json:"block_time" required:"true"`
	// Rule description.
	Description string `json:"description"`
}

// CreateKnownAttackSource will create a known attack source rule on the values in CreateKnownAttackSourceOpts.
func CreateKnownAttackSource(client *golangsdk.ServiceClient, policyId string, opts CreateKnownAttackSourceOpts) (*KnownAttackSourceRule, error) {
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

	var res KnownAttackSourceRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type KnownAttackSourceRule struct {
	// Rule ID.
	Id string `json:"id"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// Time the rule is created. The value is a 13-digit timestamp in ms.
	CreatedAt int64 `json:"timestamp"`
	// Type of the know attack source rule.
	Category string `json:"category"`
	// Rule description.
	Description string `json:"description"`
	// Block duration, in seconds.
	BlockTime int `json:"block_time"`
}
