package policies

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create will create a new backup policy based on the values in CreateOpts. To extract
// the Backup object from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToBackupPolicyCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	raw, err := client.Post(client.ServiceURL("policies"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
