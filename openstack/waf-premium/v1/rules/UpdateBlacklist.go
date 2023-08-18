package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateBlacklistOpts struct {
	// Rule name.
	Name string `json:"name"`
	// Rule description.
	Description string `json:"description"`
	// IP addresses or an IP address range.
	// IP addresses: IP addresses to be added to the blacklist or whitelist,
	// for example, 192.x.x.3
	// IP address range: IP address and subnet mask,
	// for example, 10.x.x.0/24
	Addresses string `json:"addr" required:"true"`
	// Protective action. The value can be:
	// 0: WAF blocks the requests that hit the rule.
	// 1: WAF allows the requests that hit the rule.
	// 2: WAF only logs the requests that hit the rule.
	Action string `json:"white" required:"true"`
	// ID of a known attack source rule. This parameter can be configured only when white is set to 0.
	FollowedActionId string `json:"followed_action_id"`
}

// UpdateBlacklist is used to update an IP address blacklist or whitelist rule.
func UpdateBlacklist(client *golangsdk.ServiceClient, policyId, ruleId string, opts UpdateBlacklistOpts) (*BlacklistRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/whiteblackip/{rule_id}
	raw, err := client.Put(client.ServiceURL("waf", "policy", policyId, "whiteblackip", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}

	var res BlacklistRule
	return &res, extract.Into(raw.Body, &res)
}
