package alarm

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type ListHistoryQueryOpts struct {
	// The value is the ID of the last record on the previous page
	// (value of previous_marker or next_marker returned by the previous page).
	Marker string `q:"marker"`
	// Number of records on each page. The maximum value is 100.
	Limit *int `q:"limit"`
	// Whether the alarm is an active alarm or a historical alarm.
	Type string `q:"type"`
}

type ListHistoryBodyOpts struct {
	// Whether to customize the query time segment.
	WhetherCustomField *bool `json:"whether_custom_field" required:"true"`
	// Start time of a customized time segment (timestamp).
	StartTime int64 `json:"start_time,omitempty"`
	// End time of a customized time segment (timestamp).
	EndTime int64 `json:"end_time,omitempty"`
	// Time range specified to query data of the last N
	// minutes when the client time is inconsistent with the server time.
	// It can also be used to accurately query the data of a specified period.
	TimeRange string `json:"time_range,omitempty"`
	// Field specified for fuzzy query, which can be left blank.
	// If the value is not empty, fuzzy match will be performed. The metadata field is mandatory.
	Search string `json:"search,omitempty"`
	// Alarm severity (Critical, Major, Minor, Info).
	AlarmLevelIds []string `json:"alarm_level_ids,omitempty"`
	// Sorting order, which can be left blank.
	Sort *Sort `json:"sort,omitempty"`
	// Statistical step. Unit: ms.
	// For example, if the duration is 1 minute, set this parameter to 60000.
	Step int `json:"step,omitempty"`
}

type Sort struct {
	// List of sorted fields. Fields in this list are sorted based on the specified order.
	OrderBy []string `json:"order_by" required:"true"`
	// Sorting order.
	// The value can be:
	// asc (ascending order)
	// desc (descending order)
	Order string `json:"order" required:"true"`
}

func ListHistory(client *golangsdk.ServiceClient, domainId string, opts ListHistoryQueryOpts, body ListHistoryBodyOpts) ([]Alarms, error) {
	// POST /v2/{project_id}/{domain_id}/lts/alarms/sql-alarm/query
	b, err := build.RequestBody(body, "")
	if err != nil {
		return nil, err
	}
	url, err := golangsdk.NewURLBuilder().
		WithEndpoints(domainId, "lts", "alarms", "sql-alarm", "query").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}
	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return AlarmPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Method: "post",
		Body:   b,
	}.NewAllPages()

	if err != nil {
		return nil, err
	}
	return ExtractAlarms(pages)
}

type AlarmPage struct {
	pagination.NewSinglePageBase
}

func ExtractAlarms(r pagination.NewPage) ([]Alarms, error) {
	var s struct {
		Alarms []Alarms `json:"events"`
	}
	err := extract.Into(bytes.NewReader((r.(AlarmPage)).Body), &s)
	return s.Alarms, err
}

type Alarms struct {
	// Alert details.
	Annotations *Annotations `json:"annotations"`
	// Alarm information.
	Metadata *Metadata `json:"metadata"`
	// Arrival time (timestamp).
	ArrivesAt int64 `json:"arrives_at"`
	// Alarm clearance time (timestamp).
	EndsAt int64 `json:"ends_at"`
	// Alarm ID.
	ID int `json:"id"`
	// Alarm generation time (timestamp).
	StartsAt int64 `json:"starts_at"`
	// Time when an alarm is automatically cleared (timestamp).
	Timeout int64 `json:"timeout"`
	// Alarm rule type (SQL/keyword).
	Type string `json:"type"`
}

type Annotations struct {
	// Alarm list details.
	Message string `json:"message"`
	// Log group/stream ID.
	LogInfo string `json:"log_info"`
	// Current value.
	CurrentValue string `json:"current_value"`
	// Original data of (SQL/keyword) alarm details.
	OldAnnotations string `json:"old_annotations"`
}

type Metadata struct {
	// Alarm type.
	Type string `json:"event_type"`
	// Alarm ID.
	ID string `json:"event_id"`
	// Alarm severity.
	Severity string `json:"event_severity"`
	// Alarm name.
	Name string `json:"event_name"`
	// Resource type.
	ResourceType string `json:"resource_type"`
	// Log group/stream name.
	ResourceId string `json:"resource_id"`
	// Alarm source.
	ResourceProvider string `json:"resource_provider"`
	// Alarm rule type (SQL/keyword).
	LtsAlarmType string `json:"lts_alarm_type"`
}
