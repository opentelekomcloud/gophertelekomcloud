package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BlacklistCreateOpts struct {
	// Rule name.
	Name string `json:"name"`
	// Rule description.
	Description string `json:"description"`
	// IP addresses or an IP address range.
	// IP addresses: IP addresses to be added to the blacklist or whitelist,
	// for example, 192.x.x.3 -IP address range: IP address and subnet mask, for example, 10.x.x.0/24
	Addresses string `json:"addr" required:"true"`
	// Protective action. The value can be:
	// 0: WAF blocks the requests that hit the rule.
	// 1: WAF allows the requests that hit the rule.
	// 2: WAF only logs the requests that hit the rule.
	Action *int `json:"white" required:"true"`
	// ID of a known attack source rule. This parameter can be configured only when white is set to 0.
	FollowedActionId string `json:"followed_action_id"`
}

// CreateBlacklist will create a blacklist or whitelist rule on the values in WhitelistCreateOpts.
func CreateBlacklist(client *golangsdk.ServiceClient, policyId string, opts BlacklistCreateOpts) (*BlacklistRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/policy/{policy_id}/whiteblackip
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "whiteblackip"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res BlacklistRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type BlacklistRule struct {
	// Rule ID.
	ID string `json:"id"`
	// Rule name.
	Name string `json:"name"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// Rule creation time.
	CreatedAt int64 `json:"timestamp"`
	// Rule description.
	Description string `json:"description"`
	// Rule status. The value can be:
	// 0: The rule is disabled.
	// 1: The rule is enabled.
	Status string `json:"status"`
	// Blacklisted or whitelisted IP addresses
	Addresses string `json:"addr"`
	// Protective action. The value can be:
	// 0: WAF blocks the requests that hit the rule.
	// 1: WAF allows the requests that hit the rule.
	// 2: WAF only logs the requests that hit the rule.
	Action int `json:"white"`
	// ID of the known attack source rule.
	FollowedActionId string `json:"followed_action_id"`
}
