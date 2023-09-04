package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateGeoIpOpts struct {
	// Applicable regions. The value can be the region code.
	GeoIp string `json:"geoip" required:"true"`
	// Protective action. The value can be:
	// 0: WAF blocks the requests that hit the rule.
	// 1: WAF allows the requests that hit the rule.
	// 2: WAF only logs the requests that hit the rule.
	Action *int `json:"white" required:"true"`
	// Rule name. Currently, the console does not support configuring
	// names for geolocation access control rule. Ignore this parameter.
	Name string `json:"name" required:"true"`
	// Rule description.
	Description string `json:"description"`
}

// CreateGeoIp will create a geolocation access control rule on the values in CreateOpts.
func CreateGeoIp(client *golangsdk.ServiceClient, policyId string, opts CreateGeoIpOpts) (*GeoIpRule, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v1/{project_id}/waf/policy/{policy_id}/geoip
	raw, err := client.Post(client.ServiceURL("waf", "policy", policyId, "geoip"), b,
		nil, &golangsdk.RequestOpts{
			OkCodes:     []int{200},
			MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
		})
	if err != nil {
		return nil, err
	}

	var res GeoIpRule
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type GeoIpRule struct {
	// Rule ID.
	ID string `json:"id"`
	// Policy ID.
	PolicyId string `json:"policyid"`
	// Rule name.
	Name string `json:"name"`
	// List of geographical locations hit the geolocation access control rule.
	GeoTagList []string `json:"geoTagList"`
	// Applicable regions.
	GeoIp string `json:"geoip"`
	// Protective action.
	Action int `json:"white"`
	// Rule status.
	Status int `json:"status"`
	// Time the rule is created. The value is a 13-digit timestamp in ms.
	CreatedAt int64 `json:"timestamp"`
	// Rule description.
	Description string `json:"description"`
}
