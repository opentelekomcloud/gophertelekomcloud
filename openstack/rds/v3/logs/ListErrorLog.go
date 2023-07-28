package logs

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

type DbErrorlogOpts struct {
	// Specifies the ID of the queried DB instance.
	InstanceId string `json:"-"`
	// Specifies the start time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	StartDate string `q:"start_date" required:"true"`
	// Specifies the end time in the "yyyy-mm-ddThh:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	// You can only query error logs generated within a month.
	EndDate string `q:"end_date" required:"true"`
	// Specifies the page offset, such as 1, 2, 3, or 4. The parameter value is 1 by default if it is not specified.
	Offset string `q:"offset"`
	// Specifies the number of records on a page. Its value range is from 1 to 100. The parameter value is 10 by default if it is not specified.
	Limit string `q:"limit"`
	// Specifies the log level. The default value is ALL. Valid value:
	//
	// ALL
	// INFO
	// LOG
	// WARNING
	// ERROR
	// FATAL
	// PANIC
	// NOTE
	Level string `q:"level"`
}

func ListErrorLog(client *golangsdk.ServiceClient, opts DbErrorlogOpts) (*ErrorLogResp, error) {
	query, err := build.QueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/errorlog
	url := client.ServiceURL("instances", opts.InstanceId, "errorlog") + query.String()
	raw, err := client.Get(url, nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res ErrorLogResp
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ErrorLogResp struct {
	// Indicates detailed information.
	ErrorLogList []Errorlog `json:"error_log_list"`
	// Indicates the total number of records.
	TotalRecord int `json:"total_record"`
}

type Errorlog struct {
	// Indicates the time in the UTC format.
	Time string `json:"time"`
	// Indicates the log level.
	Level string `json:"level"`
	// Indicates the log content.
	Content string `json:"content"`
}
