package monitor

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListAlarmRuleOpts struct {
	// Value range: 1-1000. Default value: 1000.
	//
	// Maximum number of returned records.
	//
	// Value Range
	//
	// (0,1000]
	Limit *int `q:"limit"`
	// Pagination information.
	//
	// Value Range
	//
	// [0, (Maximum value of the int type - 1000)]
	Start *int `q:"start"`
}

func ListAlarmRule(client *golangsdk.ServiceClient, opts ListAlarmRuleOpts) (*ListAlarmRuleResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v2/{project_id}/ams/alarms
	raw, err := client.Get(client.ServiceURL("ams", "alarms")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListAlarmRuleResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
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
