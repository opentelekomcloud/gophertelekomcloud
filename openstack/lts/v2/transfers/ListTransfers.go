package transfers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListTransfersOpts struct {
	// Log transfer type. You can transfer logs to OBS, DIS, and DMS.
	//
	// Enumerated values:
	//
	// OBS
	// DIS
	// DMS
	LogTransferType string `json:"log_transfer_type,omitempty"`
	// Log group name.
	//
	// Minimum length: 1 character
	//
	// Maximum length: 64 characters
	LogGroupName string `json:"log_group_name,omitempty"`
	// Log stream name.
	//
	// Minimum length: 1 character
	//
	// Maximum length: 64 characters
	LogStreamName string `json:"log_stream_name,omitempty"`
	// Query cursor. Set the value to 0 in the first query. In subsequent queries, obtain the value from the response to the last request.
	//
	// Minimum value: 0
	//
	// Maximum value: 1024
	Offset int32 `json:"offset,omitempty"`
	//
	// Number of records on each page.
	//
	// Minimum value: 0
	//
	// Maximum value: 100
	Limit int32 `json:"limit,omitempty"`
}

func ListTransfers(client *golangsdk.ServiceClient, opts ListTransfersOpts) ([]Transfer, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/transfers
	raw, err := client.Get(client.ServiceURL("transfers")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Transfer
	err = extract.IntoSlicePtr(raw.Body, &res, "log_transfers")
	return res, err
}

type Transfer struct {
	// Log group ID.
	//
	// Minimum length: 36 characters
	//
	// Maximum length: 36 characters
	LogGroupId string `json:"log_group_id"`
	//
	// Log group name.
	//
	// Minimum length: 1 character
	//
	// Maximum length: 64 characters
	LogGroupName string `json:"log_group_name"`
	// Log stream list.
	LogStreams []LogStreams `json:"log_streams"`
	//
	// Log transfer task ID.
	//
	// Minimum length: 36 characters
	//
	// Maximum length: 36 characters
	LogTransferId string `json:"log_transfer_id"`
	// Log transfer information.
	LogTransferInfo LogTransferInfo `json:"log_transfer_info"`
}
