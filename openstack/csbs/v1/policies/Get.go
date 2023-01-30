package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

// Get will get a single backup policy with specific ID. call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, policyId string) (*BackupPolicy, error) {
	// GET https://{endpoint}/v1/{project_id}/policies/{policy_id}
	raw, err := client.Get(client.ServiceURL("policies", policyId), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	return extra(err, raw)
}
