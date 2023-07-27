package rules

import "github.com/opentelekomcloud/gophertelekomcloud"

func Delete(client *golangsdk.ServiceClient, policyID, id string) (err error) {
	_, err = client.Delete(client.ServiceURL("l7policies", policyID, "rules", id), nil)
	return
}
