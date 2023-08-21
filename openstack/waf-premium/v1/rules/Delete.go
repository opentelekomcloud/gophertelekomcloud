package rules

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// DeleteBlacklistRule is used to delete an IP address blacklist or whitelist rule.
func DeleteBlacklistRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/whiteblackip/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "whiteblackip", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeleteCcRule is used to delete a CC attack protection rule.
func DeleteCcRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/cc/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "cc", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeleteCustomRule is used to delete a precise protection rule.
func DeleteCustomRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/custom/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "custom", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeleteAntiCrawlerRule is used to delete a JavaScript anti-crawler rule.
func DeleteAntiCrawlerRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/anticrawler/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "anticrawler", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeletePrivacyRule is used to delete a data masking rule.
func DeletePrivacyRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/privacy/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "privacy", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeleteKnownAttackSourceRule is used to delete a known attack source rule.
func DeleteKnownAttackSourceRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/punishment/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "punishment", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeleteAntiTamperRule is used to delete a web tamper protection rule.
func DeleteAntiTamperRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/antitamper/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "antitamper", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeleteAntiLeakageRule is used to delete an information leakage prevention rule.
func DeleteAntiLeakageRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/antileakage/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "antileakage", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeleteIgnoreRule is used to deleting a global protection whitelist (false alarm masking) rule.
func DeleteIgnoreRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/ignore/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "ignore", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeleteGeoIpRule is used to delete a geolocation access control rule.
func DeleteGeoIpRule(client *golangsdk.ServiceClient, policyId, ruleId string) (err error) {
	// DELETE /v1/{project_id}/waf/policy/{policy_id}/geoip/{rule_id}
	_, err = client.Delete(client.ServiceURL("waf", "policy", policyId, "geoip", ruleId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}

// DeleteReferenceTable is used to delete a reference table.
func DeleteReferenceTable(client *golangsdk.ServiceClient, tableId string) (err error) {
	// DELETE /v1/{project_id}/waf/valuelist/{table_id}
	_, err = client.Delete(client.ServiceURL("waf", "valuelist", tableId), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json;charset=utf8"},
	})
	return
}
