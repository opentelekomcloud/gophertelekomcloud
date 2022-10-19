package backup

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete will delete an existing backup.
func Delete(client *golangsdk.ServiceClient, checkpointID string) (err error) {
	// https://{endpoint}/v1/{project_id}/providers/{provider_id}/checkpoints/{checkpoint_id}
	_, err = client.Delete(client.ServiceURL("providers", ProviderID, "checkpoints", checkpointID), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
