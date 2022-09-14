package backup

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will delete an existing backup.
func Delete(client *golangsdk.ServiceClient, checkpointID string) (r DeleteResult) {
	_, r.Err = client.Delete(client.ServiceURL("providers", providerID, "checkpoints", checkpointID), &golangsdk.RequestOpts{
		OkCodes:      []int{200},
		JSONResponse: nil,
	})
	return
}
