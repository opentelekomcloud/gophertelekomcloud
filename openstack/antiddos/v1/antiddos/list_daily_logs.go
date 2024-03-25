package antiddos

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListDailyLogsOps struct {
	// Limit of number of returned results or the maximum number of returned results of a query.
	// The value ranges from 1 to 100, and this parameter is used together with the offset parameter.
	// If neither limit nor offset is used, query results of all ECSs are returned.
	Limit int `q:"limit"`
	// This parameter is valid only when used together with the limit parameter.
	Offset int `q:"offset"`
	// Possible values:
	// desc: indicates that query results are given and sorted by time in descending order.
	// asc: indicates that query results are given and sorted by time in ascending order.
	// The default value is desc.
	SortDir string `q:"sort_dir"`
}

func ListDailyLogs(client *golangsdk.ServiceClient, floatingIpId string, opts ListDailyLogsOps) (*ListDailyLogsResponse, error) {
	query, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET /v1/{project_id}/antiddos/{floating_ip_id}/logs
	raw, err := client.Get(client.ServiceURL("antiddos", floatingIpId, "logs")+query.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListDailyLogsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListDailyLogsResponse struct {
	// Total number of EIPs
	Total int `json:"total"`
	// List of events
	Logs []Logs `json:"logs"`
}

type Logs struct {
	// Start time
	StartTime int `json:"start_time"`
	// End time
	EndTime int `json:"end_time"`
	// Defense status, the possible value of which is one of the following:
	// 1: indicates that traffic cleaning is underway.
	// 2: indicates that traffic is discarded.
	Status int `json:"status"`
	// Traffic at the triggering point.
	TriggerBps int `json:"trigger_bps"`
	// Packet rate at the triggering point
	TriggerPps int `json:"trigger_pps"`
	// HTTP request rate at the triggering point
	TriggerHttpPps int `json:"trigger_http_pps"`
}
