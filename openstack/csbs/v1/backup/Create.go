package backup

import "github.com/opentelekomcloud/gophertelekomcloud"

// Create will create a new backup based on the values in CreateOpts. To extract
// the checkpoint object from the response, call the Extract method on the CreateResult.
func Create(client *golangsdk.ServiceClient, resourceID string, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToBackupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(client.ServiceURL("providers", providerID, "resources", resourceID, "action"), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
