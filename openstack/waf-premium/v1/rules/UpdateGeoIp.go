package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateGeoIpOpts struct {
	// Applicable regions.
	GeoIp string `json:"geoip" required:"true"`
	// Protective action. The value can be:
	// 0: WAF blocks the requests that hit the rule.
	// 1: WAF allows the requests that hit the rule.
	// 2: WAF only logs the requests that hit the rule.
	Action int `json:"white" required:"true"`
	// Name of the masked field
	Name string `json:"name"`
	// Rule description
	Description string `json:"description"`
}

// UpdateGeoIp is used to update a geolocation access control rule.
func UpdateGeoIp(client *golangsdk.ServiceClient, policyId, ruleId string, opts UpdateGeoIpOpts) (*GeoIpRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v1/{project_id}/waf/policy/{policy_id}/geoip/{rule_id}
	raw, err := client.Put(client.ServiceURL("waf", "policy", policyId, "geoip", ruleId), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	if err != nil {
		return nil, err
	}
	var res GeoIpRule
	return &res, extract.Into(raw.Body, &res)
}
