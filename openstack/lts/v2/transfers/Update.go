package transfers

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateTransferOpts struct {
	TransferId   string              `json:"log_transfer_id" required:"true"`
	TransferInfo *TransferInfoUpdate `json:"log_transfer_info" required:"true"`
}

type TransferInfoUpdate struct {
	// Log transfer format. The value can be RAW or JSON.
	// RAW indicates raw log format, whereas JSON indicates JSON format.
	// JSON and RAW are supported for OBS and DIS transfer tasks, but only RAW is supported for DMS transfer tasks.
	// Enumerated values:
	// JSON
	// RAW
	StorageFormat string `json:"log_storage_format" required:"true"`
	// Log transfer status. ENABLE indicates that log transfer is enabled, DISABLE indicates that log transfer is disabled, and EXCEPTION indicates that log transfer is abnormal.
	// Enumerated values:
	// ENABLE
	// DISABLE
	// EXCEPTION
	TransferStatus string `json:"log_transfer_status" required:"true"`
	// Log transfer details.
	TransferDetail *TransferDetail `json:"log_transfer_detail" required:"true"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateTransferOpts) (*TransferResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v2/{project_id}/transfers
	raw, err := client.Put(client.ServiceURL("transfers"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res TransferResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
