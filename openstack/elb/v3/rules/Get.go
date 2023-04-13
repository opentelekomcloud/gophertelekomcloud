package rules

import "github.com/opentelekomcloud/gophertelekomcloud"

func Get(client *golangsdk.ServiceClient, policyID, id string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("l7policies", policyID, "rules", id), &r.Body, nil)
	return
}
