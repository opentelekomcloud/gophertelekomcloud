package rules

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

const (
	policyBasePath = "l7policies"
	ruleBasePath   = "rules"
)

func baseURL(client *golangsdk.ServiceClient, policyID string) string {
	return client.ServiceURL(policyBasePath, policyID, ruleBasePath)
}

func resourceURL(client *golangsdk.ServiceClient, policyID, id string) string {
	return client.ServiceURL(policyBasePath, policyID, ruleBasePath, id)
}
