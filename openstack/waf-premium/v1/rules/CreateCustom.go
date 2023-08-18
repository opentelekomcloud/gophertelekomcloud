package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateCustomOpts struct {
	// Time the precise protection rule takes effect.
	// false: The rule takes effect immediately.
	// true: The effective time is customized.
	Time *bool `json:"time" required:"true"`
	// Timestamp (ms) when the precise protection rule takes effect.
	// This parameter is returned only when time is true.
	Start int64 `json:"start"`
	// Timestamp (ms) when the precise protection rule expires.
	// This parameter is returned only when time is true.
	Terminal int64 `json:"terminal"`
	// Rule description.
	Description string `json:"description"`
	// Match condition List.
	Conditions []CustomConditionsObject `json:"conditions"`
	// Protective action of the precise protection rule.
	Action *CustomActionObject `json:"action" required:"true"`
	// Priority of a rule. A small value indicates a high priority. If two rules are assigned with the same priority,
	// the rule added earlier has higher priority. Value range: 0 to 1000.
	Priority int `json:"priority" required:"true"`
}

type CustomConditionsObject struct {
	// Field type. The value can be url, ip, params, cookie, or header.
	Category string `json:"category"`
	// Logic for matching the condition.
	// If the category is url, the optional operations are:
	// `contain`, `not_contain`, `equal`, `not_equal`, `prefix`, `not_prefix`, `suffix`, `not_suffix`,
	// `contain_any`, `not_contain_all`, `equal_any`, `not_equal_all`, `equal_any`,
	// `not_equal_all`, `prefix_any`, `not_prefix_all`, `suffix_any`, `not_suffix_all`,
	// `len_greater`, `len_less`, `len_equal` and `len_not_equal`
	// If the category is ip, the optional operations are:
	// `equal`, `not_equal`, `equal_any` and `not_equal_all`
	// If the category is params, cookie and header, the optional operations are:
	// `contain`, `not_contain`, `equal`, `not_equal`, `prefix`, `not_prefix`, `suffix`, `not_suffix`,
	// `contain_any`, `not_contain_all`, `equal_any`, `not_equal_all`, `equal_any`, `not_equal_all`,
	// `prefix_any`, `not_prefix_all`, `suffix_any`, `not_suffix_all`, `len_greater`, `len_less`,
	// `len_equal`, `len_not_equal`, `num_greater`, `num_less`, `num_equal`, `num_not_equal`,
	// `exist` and `not_exist`
	LogicOperation string `json:"logic_operation"`
	// Content of the conditions.
	// This parameter is mandatory when the suffix of logic_operation is not any or all.
	Contents []string `json:"contents"`
	// Reference table ID. It can be obtained by calling the API Querying the Reference Table List.
	// This parameter is mandatory when the suffix of logic_operation is any or all.
	// The reference table type must be the same as the category type.
	ValueListId string `json:"value_list_id"`
	// Subfield. When category is set to params, cookie, or header,
	// set this parameter based on site requirements.
	// This parameter is mandatory.
	Index string `json:"index"`
}

type CustomActionObject struct {
	// Operation type
	// block: WAF blocks attacks.
	// pass: WAF allows requests.
	// log: WAF only logs detected attacks.
	Category string `json:"category" required:"true"`
	// ID of a known attack source rule.
	// This parameter can be configured only when category is set to block.
	FollowedActionId string `json:"followed_action_id"`
}

// CreateCustom will  create a precise protection rule on the values in CreateOpts.
func CreateCustom(client *golangsdk.ServiceClient, policyId string, opts CreateCustomOpts) (*CustomRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/policy/{policy_id}/custom
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "custom"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res CustomRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type CustomRule struct {
	// Rule ID.
	Id string `json:"id"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// Rule description.
	Description string `json:"description"`
	// Rule status. The value can be 0 or 1.
	Status string `json:"status"`
	// List of matching conditions. All conditions must be met.
	Conditions []CustomConditionsObject `json:"conditions"`
	// Protective action of the precise protection rule.
	Action []CustomActionObject `json:"action"`
	// Priority of a rule. A small value indicates a high priority.
	// If two rules are assigned with the same priority,
	// the rule added earlier has higher priority. Value range: 0 to 1000.
	Priority int `json:"priority"`
	// Timestamp when the precise protection rule is created.
	CreatedAt int64 `json:"timestamp"`
	// Timestamp (ms) when the precise protection rule takes effect.
	// This parameter is returned only when time is true.
	Start int64 `json:"start"`
	// Timestamp (ms) when the precise protection rule expires.
	// This parameter is returned only when time is true.
	Terminal int64 `json:"terminal"`
	// This parameter is reserved and can be ignored currently.
	ActionMode *bool `json:"action_mode"`
	// Rule aging time. This parameter is reserved and can be ignored currently.
	AgingTime int `json:"aging_time"`
	// Rule creation object. This parameter is reserved and can be ignored currently.
	Producer int `json:"producer"`
}
