package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateAntiLeakageOpts struct {
	// URL to which the rule applies.
	Url string `json:"url" required:"true"`
	// Sensitive information type in the information leakage prevention rule.
	// sensitive: The rule masks sensitive user information, such as ID code, phone numbers,
	// and email addresses.
	// code: The rule blocks response pages of specified HTTP response code.
	Category string `json:"category" required:"true"`
	// Content corresponding to the sensitive information type. Multiple options can be set.
	// When category is set to code, the pages that contain the following HTTP response codes
	// will be blocked: 400, 401, 402, 403, 404, 405, 500, 501, 502, 503, 504 and 507.
	// When category is set to sensitive, parameters phone, id_card, and email can be set.
	Contents []string `json:"contents" required:"true"`
	// Rule description.
	Description string `json:"description"`
}

// CreateAntiLeakage will create an information leakage protection rule on the values in CreateOpts.
func CreateAntiLeakage(client *golangsdk.ServiceClient, policyId string, opts CreateAntiLeakageOpts) (*AntiLeakageRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/policy/{policy_id}/antileakage
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "antileakage"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res AntiLeakageRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type AntiLeakageRule struct {
	// Rule ID.
	Id string `json:"id"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// URL to which the rule applies.
	Url string `json:"url"`
	// Sensitive information type in the information leakage prevention rule.
	// sensitive: The rule masks sensitive user information, such as ID code,
	// phone numbers, and email addresses.
	// code: The rule blocks response pages of specified HTTP response code.
	Category string `json:"category"`
	// Content corresponding to the sensitive information type.
	Contents []string `json:"contents"`
	// Time the rule is created. The value is a 13-digit timestamp in ms.
	CreatedAt int64 `json:"timestamp"`
	// Rule status. The value can be:
	// 0: The rule is disabled.
	// 1: The rule is enabled.
	Status int `json:"status"`
	// Rule description.
	Description string `json:"description"`
}
