package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateAntiLeakageOpts struct {
	// URL to which the rule applies.
	Url string `json:"url" required:"true"`
	// Sensitive information type in the information leakage prevention rule.
	// sensitive: The rule masks sensitive user information, such as ID code, phone numbers, and email addresses.
	// code: The rule blocks response pages of specified HTTP response code.
	Category string `json:"category" required:"true"`
	// Content corresponding to the sensitive information type. Multiple options can be set.
	// When category is set to code, the pages that contain the following HTTP response codes
	// will be blocked: 400, 401, 402, 403, 404, 405, 500, 501, 502, 503, 504 and 507.
	// When category is set to sensitive, parameters phone, id_card, and email can be set.
	Contents []string `json:"contents" required:"true"`
	// Rule description
	Description string `json:"description"`
}

// UpdateAntiLeakage is used to update an information leakage prevention rule.
func UpdateAntiLeakage(client *golangsdk.ServiceClient, policyId, ruleId string, opts UpdateAntiLeakageOpts) (*AntiLeakageRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/antileakage/{rule_id}
	raw, err := client.Put(client.ServiceURL("waf", "policy", policyId, "antileakage", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res AntiLeakageRule
	return &res, extract.Into(raw.Body, &res)
}
