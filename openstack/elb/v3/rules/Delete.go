package rules

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, policyID, id string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("l7policies", policyID, "rules", id), nil)
	return
}
