package monitor

import "github.com/opentelekomcloud/gophertelekomcloud"

type ListAlarmRuleOpts struct {
	// Value range: 1-1000. Default value: 1000.
	//
	// Maximum number of returned records.
	//
	// Value Range
	//
	// (0,1000]
	Limit *int `json:"limit,omitempty"`
	// Pagination information.
	//
	// Value Range
	//
	// [0, (Maximum value of the int type - 1000)]
	Start *int `json:"start,omitempty"`
}

func ListAlarmRule(client *golangsdk.ServiceClient, opts ListAlarmRuleOpts) (*ListAlarmRuleResponse, error) {
	// GET /v2/{project_id}/ams/alarms
}

type ListAlarmRuleResponse struct {
	// Response code.
	ErrorCode string `json:"errorCode"`
	// Response message.
	ErrorMessage string `json:"errorMessage"`
	// Metadata, including pagination information.
	MetaData MetaData `json:"metaData"`
	// Threshold rule list.
	Thresholds []Threshold `json:"thresholds"`
}

type MetaData struct {
	// Number of returned records.
	Count *int `json:"count"`
	// Total number of records.
	Start string `json:"start"`
	// Start of the next page, which is used for pagination.
	Total *int `json:"total"`
}
