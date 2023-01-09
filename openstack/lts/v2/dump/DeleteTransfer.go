package dump

import "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteTransfer(client *golangsdk.ServiceClient, transferId string) (err error) {
	// DELETE /v2/{project_id}/transfers
	_, err = client.Delete(client.ServiceURL("transfers")+"?log_transfer_id="+transferId, nil)
	return
}
