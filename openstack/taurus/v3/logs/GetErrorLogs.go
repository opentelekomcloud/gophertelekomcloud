package logs

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type GetErrorLogsOpts struct {
	InstanceId string `json:"-"`
	StartDate  string `q:"start_date" required:"true"`
	EndDate    string `q:"end_date" required:"true"`
	NodeId     string `q:"node_id" required:"true"`
	Offset     int    `q:"offset"`
	Limit      int    `q:"limit"`
	Level      string `q:"level"`
}

func GetErrorLogs(client *golangsdk.ServiceClient, opts GetErrorLogsOpts) ([]MysqlErrorLogList, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("instances", opts.InstanceId, "errorlog").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return ErrorLogPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractErrorLogs(pages)
}

func ExtractErrorLogs(r pagination.NewPage) ([]MysqlErrorLogList, error) {
	var s struct {
		ErrorLogList []MysqlErrorLogList `json:"error_log_list"`
	}
	err := extract.Into(bytes.NewReader((r.(ErrorLogPage)).Body), &s)
	return s.ErrorLogList, err
}

type ErrorLogPage struct {
	pagination.NewSinglePageBase
}

type MysqlErrorLogList struct {
	NodeId  string `json:"node_id"`
	Time    string `json:"time"`
	Level   string `json:"level"`
	Content string `json:"content"`
}
