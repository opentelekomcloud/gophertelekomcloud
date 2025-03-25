package transfers

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListTransfersOpts struct {
	// Log transfer type. You can transfer logs to OBS, DIS, and DMS.
	// Enumerated values:
	// OBS
	// DIS
	// DMS
	LogTransferType string `json:"log_transfer_type,omitempty"`
	// Log group name.
	LogGroupName string `json:"log_group_name,omitempty"`
	// Log stream name.
	LogStreamName string `json:"log_stream_name,omitempty"`
	// Query cursor. Set the value to 0 in the first query.
	// In subsequent queries, obtain the value from the response to the last request.
	// Minimum value: 0
	// Maximum value: 1024
	Offset int32 `json:"offset,omitempty"`
	// Number of records on each page.
	// Minimum value: 0
	// Maximum value: 100
	Limit int32 `json:"limit,omitempty"`
}

func List(client *golangsdk.ServiceClient, opts ListTransfersOpts) ([]Transfer, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("transfers").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/transfers
	raw, err := client.Get(client.ServiceURL(url.String()), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res []Transfer
	err = extract.IntoSlicePtr(raw.Body, &res, "log_transfers")
	return res, err
}

type Transfer struct {
	// Log group ID.
	LogGroupId string `json:"log_group_id"`
	// Log group name.
	LogGroupName string `json:"log_group_name"`
	// Log stream list.
	LogStreams []LogStreamsResponse `json:"log_streams"`
	// Log transfer task ID.
	LogTransferId string `json:"log_transfer_id"`
	// Log transfer information.
	LogTransferInfo LogTransferInfoResponse `json:"log_transfer_info"`
}
