package v3

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListSlowLogsOpts struct {
	// Instance ID
	InstanceId string `q:"instance_id"`
	// Start time in the "yyyy-mm-ddThh:mm:ssZ" format.
	StartDate string `q:"start_date"`
	// End time in the "yyyy-mm-ddThh:mm:ssZ" format.
	EndDate string `q:"end_date"`
	// Node ID. If this parameter is left blank, all nodes of the instance are queried.
	NodeId string `q:"node_id,omitempty"`
	// Statement type. If this parameter is left empty, all statement types are queried.
	// You can also specify the following log type:
	// SELECT
	Type string `q:"type,omitempty"`
	// Index offset. Its value range is [0, 1999].
	// If offset is set to N, the resource query starts from the N+1 piece of data. The value is 0 by default,
	// indicating that the query starts from the first piece of data. The value must be a positive number.
	Offset int32 `q:"offset,omitempty"`
	// Number of records to be queried. The value ranges from 1 (inclusive) to 100 (inclusive).
	// The sum of values of limit and offset must be less than or equal to 2000.
	Limit int32 `q:"limit,omitempty"`
}

func ListSlowLogs(client *golangsdk.ServiceClient, opts ListSlowLogsOpts) (*ListSlowLogsResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/slowlog
	raw, err := client.Get(client.ServiceURL("instances", opts.InstanceId, "slowlog")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListSlowLogsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListSlowLogsResponse struct {
	SlowLogList []SlowlogResult `json:"slow_log_listq"`
	TotalRecord int32           `json:"total_recordq"`
}

type SlowlogResult struct {
	// Execution time
	Time string `json:"time"`
	// Database which slow logs belong to
	Database string `json:"database"`
	// Execution syntax
	QuerySample string `json:"query_sample"`
	// Statement type
	Type string `json:"type"`
	// Time in the UTC format
	StartTime string `json:"start_time"`
}
