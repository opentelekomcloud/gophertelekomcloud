package flow_logs

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListOpts struct {
	// ID of the last enterprise router on the previous page.
	// If this parameter is left blank, the first page is queried.
	// This parameter must be used together with limit.
	Marker string `q:"marker"`
	// Number of records on each page. Value range: 0 to 2000
	Limit int `q:"limit"`
	// Resource type
	ResourceType string `q:"resource_type"`
	// Attachment resource IDs
	ResourceID []string `q:"resource_id"`
	// Keyword for sorting. The keyword can be id, name, or state. By default, id is used.
	SortKey []string `q:"sort_key"`
	// Sorting order. There are two value options: asc (ascending order) and desc (descending order).
	// The default value is asc.
	SortDir []string `q:"sort_dir"`
}

func List(client *golangsdk.ServiceClient, routerID string, opts ListOpts) ([]FlowLogResponse, error) {
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("enterprise-router", routerID, "flow-logs").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return FlowLogPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractFlowLogs(pages)
}

type FlowLogPage struct {
	pagination.NewSinglePageBase
}

func ExtractFlowLogs(r pagination.NewPage) ([]FlowLogResponse, error) {
	var s struct {
		FlowLogs []FlowLogResponse `json:"flow_logs"`
	}
	err := extract.Into(bytes.NewReader((r.(FlowLogPage)).Body), &s)
	return s.FlowLogs, err
}
