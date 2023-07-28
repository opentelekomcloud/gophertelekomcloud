package logs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type DbSlowLogOpts struct {
	// Specifies the ID of the queried DB instance.
	InstanceId string `json:"-"`
	// Specifies the start date in the "yyyy-mm-ddThh:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	StartDate string `q:"start_date" required:"true"`
	// Specifies the end time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset. You can only query slow logs generated within a month.
	EndDate string `q:"end_date" required:"true"`
	// Specifies the page offset, such as 1, 2, 3, or 4. The parameter value is 1 by default if it is not specified.
	Offset string `q:"offset"`
	// Specifies the number of records on a page. Its value range is from 1 to 100. The parameter value is 10 by default if it is not specified.
	Limit string `q:"limit"`
	// Specifies the statement type. If it is left blank, all statement types are queried. Valid value:
	//
	// INSERT
	// UPDATE
	// SELECT
	// DELETE
	// CREATE
	Level string `q:"level"`
}

func ListSlowLog(client *golangsdk.ServiceClient, opts DbSlowLogOpts) (*SlowLogResp, error) {
	var opts2 interface{} = opts
	query, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/slowlog
	url := client.ServiceURL("instances", opts.InstanceId, "slowlog") + query.String()
	raw, err := client.Get(url, nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res SlowLogResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type SlowLogResp struct {
	// Indicates detailed information.
	Slowloglist []Slowloglist `json:"slow_log_list"`
	// Indicates the total number of records.
	TotalRecord int `json:"total_record"`
}

type Slowloglist struct {
	// Indicates the number of executions.
	Count string `json:"count"`
	// Indicates the execution time.
	Time string `json:"time"`
	// Indicates the lock wait time.
	// This parameter is not present in the response for PostgreSQL DB engine.
	LockTime string `json:"lock_time"`
	// Indicates the number of sent rows.
	// This parameter is not present in the response for PostgreSQL DB engine.
	RowsSent string `json:"rows_sent"`
	// Indicates the number of scanned rows.
	// This parameter is not present in the response for PostgreSQL DB engine.
	RowsExamined string `json:"rows_examined"`
	// Indicates the database which the slow log belongs to.
	Database string `json:"database"`
	// Indicates the account.
	Users string `json:"users"`
	// Indicates the execution syntax.
	QuerySample string `json:"query_sample"`
	// Indicates the statement type.
	Type string `json:"type"`
	// Indicates the time in the UTC format.
	StartTime string `json:"start_time"`
	// Indicates the IP address.
	ClientIp string `json:"client_ip"`
}
