package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateAntiTamperOpts struct {
	// Protected website. It can be obtained by calling the ListHost API
	// in cloud mode (the value of the hostname field in the response body).
	Hostname string `json:"hostname" required:"true"`
	// URL protected by the web tamper protection rule.
	// The value must be in the standard URL format, for example, /admin
	Url string `json:"url" required:"true"`
	// Rule description.
	Description string `json:"description"`
}

// CreateAntiTamper will create a web tamper protection rule on the values in CreateAntiTamperOpts.
func CreateAntiTamper(client *golangsdk.ServiceClient, policyId string, opts CreateAntiTamperOpts) (*AntiTamperRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/policy/{policy_id}/antitamper
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "antitamper"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res AntiTamperRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type AntiTamperRule struct {
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
	// Rule description.
	Status *int `json:"status"`
	// The domain name of the website protected with the web tamper protection rule.
	// The domain name is in the format of xxx.xxx.com, such as www.example.com.
	Hostname string `json:"hostname"`
	// URL for the web tamper protection rule.
	Url string `json:"url"`
}
