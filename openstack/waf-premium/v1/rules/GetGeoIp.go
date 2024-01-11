package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// GetGeoIp is used to query a geolocation access control rule by ID.
func GetGeoIp(client *golangsdk.ServiceClient, policyId, ruleId string) (*GeoIpRule, error) {
	// GET /v1/{project_id}/waf/policy/{policy_id}/geoip/{rule_id}
	raw, err := client.Get(client.ServiceURL("waf", "policy", policyId, "geoip", ruleId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GeoIpRule
	return &res, extract.Into(raw.Body, &res)
}
