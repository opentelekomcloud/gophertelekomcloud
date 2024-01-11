package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateIgnoreOpts struct {
	// Domain names to be protected. If the array length is 0, this rule will take effect
	// for all domain names that are protected by the policies this rule belongs to.
	Domains []string `json:"domain" required:"true"`
	// Condition list
	Conditions []IgnoreCondition `json:"conditions" required:"true"`
	// The value is fixed at 1, indicating v2 false alarm masking rules.
	Mode int `json:"mode" required:"true"`
	// Items to be masked. You can provide multiple items and separate them with semicolons (;).
	Rule string `json:"rule" required:"true"`
	// To ignore attacks of a specific field, specify the field in the Advanced settings area.
	// After you add the rule, WAF will stop blocking attacks of the specified field.
	// This parameter is not included if all modules are bypassed.
	Advanced []AdvancedIgnoreObject `json:"advanced"`
	// Description of the rule
	Description string `json:"description,omitempty"`
}

type IgnoreCondition struct {
	// Field type. The value can be url, ip, params, cookie, or header.
	Category string `json:"category,omitempty"`
	// Content. The array length is limited to 1.
	// The content format varies depending on the field type.
	// For example, if the field type is ip, the value must be an IP address or IP address range.
	// If the field type is url, the value must be in the standard URL format.
	// IF the field type is params, cookie, or header, the content format is not limited.
	Contents []string `json:"contents,omitempty"`
	// The matching logic varies depending on the field type. For example,
	// if the field type is ip, the logic can be equal or not_equal.
	// If the field type is url, params, cookie, or header,
	// the logic can be equal, not_equal, contain, not_contain, prefix, not_prefix,
	// suffix, not_suffix.
	LogicOperation string `json:"logic_operation,omitempty"`
	// If the field type is ip and the subfield is the client IP address,
	// the index parameter is not required. If the subfield type is X-Forwarded-For,
	// the value is x-forwarded-for; If the field type is params, header,
	// or cookie, and the subfield is user-defined, the value of index is the user-defined subfield.
	Index string `json:"index,omitempty"`
}

type AdvancedIgnoreObject struct {
	// Field type. The following field types are supported: Params, Cookie, Header, Body, and Multipart.
	// When you select Params, Cookie, or Header, you can set this parameter to all or configure subfields as required.
	// When you select Body or Multipart, set this parameter to all.
	Index string `json:"index,omitempty"`
	// Subfield of the specified field type. The default value is all.
	Contents []string `json:"contents,omitempty"`
}

// CreateIgnore will create a global protection whitelist (formerly false alarm masking) rule on the values in CreateOpts.
func CreateIgnore(client *golangsdk.ServiceClient, policyId string, opts CreateIgnoreOpts) (*IgnoreRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/policy/{policy_id}/ignore
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "ignore"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res IgnoreRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type IgnoreRule struct {
	// Rule ID.
	ID string `json:"id"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// Time the rule is created. The value is a 13-digit timestamp in ms.
	CreatedAt int64 `json:"timestamp"`
	// Rule description.
	Description string `json:"description"`
	// Rule status. The value can be:
	// 0: The rule is disabled.
	// 1: The rule is enabled.
	Status int `json:"status"`
	// Masked items.
	Rule string `json:"rule"`
	// The value is fixed at 1, indicating v2 false alarm masking rules are used.
	Mode int `json:"mode"`
	// Condition list.
	Conditions []IgnoreCondition `json:"conditions"`
	// Advanced settings.
	Advanced []AdvancedIgnoreObject `json:"advanced"`
	// Domain names.
	Domains []string `json:"domain"`
}
