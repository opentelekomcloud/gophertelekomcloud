package volumetransfers

import "github.com/opentelekomcloud/gophertelekomcloud"

// Delete deletes a volume transfer.
func Delete(client *golangsdk.ServiceClient, id string) (err error) {
	// DELETE /v3/{project_id}/os-volume-transfer/{transfer_id}
	_, err = client.Delete(client.ServiceURL("os-volume-transfer", id), nil)
	return
}
