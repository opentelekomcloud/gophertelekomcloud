package streams

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListLogsOpts struct {
	GroupId  string `json:"-" required:"true"`
	StreamId string `json:"-" required:"true"`
	// UTC start time of the search window (in milliseconds).
	// NOTE:
	// Maximum query time range: 30 days
	StartTime string `json:"start_time" required:"true"`
	// UTC end time of the search window (in milliseconds).
	// NOTE:
	// Maximum query time range: 30 days
	EndTime string `json:"end_time" required:"true"`
	// Filter criteria, which vary between log sources.
	Labels map[string]string `json:"labels,omitempty"`
	// Keyword used for search. A keyword is a word between two adjacent delimiters.
	// Enumerated value:
	// error
	Keywords string `json:"keywords,omitempty"`
	// Sequence number of a log event. This parameter is not required for the first query, but is required for subsequent pagination queries. The value can be obtained from the response of the last query. The value of line_num should be between the values of start_time and end_time.
	// Value length: 19 characters
	LineNum string `json:"line_num,omitempty"`
	// Whether the search order is descending or ascending. The default value is false, indicating that search results are displayed in ascending order.
	// Enumerated value:
	// false
	IsDesc *bool `json:"is_desc,omitempty"`
	// The value is init (default value) for the first query, or forwards or backwards for a pagination query. This parameter is used together with is_desc for pagination queries.
	// Enumerated value:
	// forwards
	SearchType string `json:"search_type,omitempty"`
	// Number of logs to be queried each time. The value is 50 when this parameter is not set. You are advised to set this parameter to 100.
	//
	// Minimum value: 1
	// Maximum value: 5000
	Limit int32 `json:"limit,omitempty"`
}

func ListLogs(client *golangsdk.ServiceClient, opts ListLogsOpts) (*ListLogsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/groups/{log_group_id}/streams/{log_stream_id}/content/query
	raw, err := client.Post(client.ServiceURL("groups", opts.GroupId, "streams", opts.StreamId, "content", "query"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res ListLogsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListLogsResponse struct {
	Count int           `json:"count,omitempty"`
	Logs  []LogContents `json:"logs,omitempty"`
}

type LogContents struct {
	// Raw log data.
	// Minimum length: 1 character
	// Maximum length: 10,000 characters
	Content string `json:"content,omitempty"`
	// Sequence number of a log event.
	// Value length: 19 characters
	LineNum string `json:"line_num,omitempty"`
	// Labels contained in a log event. The labels vary depending on log events.
	Labels map[string]string `json:"labels,omitempty"`
}
