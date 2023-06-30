package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will delete an existing backup policy.
func Delete(client *golangsdk.ServiceClient, policyId string) (err error) {
	_, err = client.Delete(client.ServiceURL("policies", policyId), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
