package transfers

import "github.com/opentelekomcloud/gophertelekomcloud"

func DeleteTransfer(client *golangsdk.ServiceClient, transferId string) (err error) {
	// DELETE /v2/{project_id}/transfers
	_, err = client.Delete(client.ServiceURL("transfers")+"?log_transfer_id="+transferId, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return
}
