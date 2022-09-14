package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get will get a single backup policy with specific ID.
// call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, policyId string) (r GetResult) {
	_, r.Err = client.Get(client.ServiceURL("policies", policyId), &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
