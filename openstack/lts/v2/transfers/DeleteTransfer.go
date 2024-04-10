package transfers

import "github.com/opentelekomcloud/gophertelekomcloud"

type deleteOpts struct {
	LogTransferId string `q:"log_transfer_id"`
}

func DeleteTransfer(client *golangsdk.ServiceClient, transferId string) (err error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("transfers").WithQueryParams(&deleteOpts{LogTransferId: transferId}).Build()
	if err != nil {
		return err
	}

	// DELETE /v2/{project_id}/transfers
	_, err = client.Delete(client.ServiceURL(url.String()), &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	return
}
