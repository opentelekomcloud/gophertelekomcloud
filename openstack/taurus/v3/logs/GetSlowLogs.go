package logs

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type GetSlowLogsOpts struct {
	InstanceId string `q:"-"`
	StartDate  string `q:"start_date" required:"true"`
	EndDate    string `q:"end_date" required:"true"`
	NodeId     string `q:"node_id" required:"true"`
	Offset     int    `q:"offset"`
	Limit      int    `q:"limit"`
	Type       string `q:"type"`
}

func GetSlowLogs(client *golangsdk.ServiceClient, opts GetSlowLogsOpts) ([]MysqlSlowLogList, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", opts.InstanceId, "slowlog").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return SlowLogPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractSlowLogs(pages)
}

func ExtractSlowLogs(r pagination.NewPage) ([]MysqlSlowLogList, error) {
	var s struct {
		SlowLogList []MysqlSlowLogList `json:"slow_log_list"`
	}
	err := extract.Into(bytes.NewReader((r.(SlowLogPage)).Body), &s)
	return s.SlowLogList, err
}

type SlowLogPage struct {
	pagination.NewSinglePageBase
}

type MysqlSlowLogList struct {
	NodeId       string `json:"node_id"`
	Count        string `json:"count"`
	Time         string `json:"time"`
	LockTime     string `json:"lock_time"`
	RowsSent     string `json:"rows_sent"`
	RowsExamined string `json:"rows_examined"`
	Database     string `json:"database"`
	Users        string `json:"users"`
	QuerySample  string `json:"query_sample"`
	Type         string `json:"type"`
	StartTime    string `json:"start_time"`
	ClientIp     string `json:"client_ip"`
}
