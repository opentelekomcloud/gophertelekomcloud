package log_statistics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListTopNTrafficStatisticsOpts struct {
	// EndTime is the end timestamp in milliseconds.
	EndTime *int64 `json:"end_time" required:"true"`
	// IsDesc controls whether data is sorted in descending order.
	IsDesc *bool `json:"is_desc" required:"true"`
	// ResourceType is log_group, log_stream, or tenant.
	ResourceType string `json:"resource_type" required:"true"`
	// SortBy is index, write, or storage.
	SortBy string `json:"sort_by" required:"true"`
	// StartTime is the start timestamp in milliseconds.
	StartTime *int64 `json:"start_time" required:"true"`
	// TopN is the number of records to return, from 1 to 100.
	TopN int `json:"topn" required:"true"`
	// Filter contains exact-match log_group_id or log_stream_id filters.
	Filter *map[string]string `json:"filter" required:"true"`
	// SearchList contains one or more of index, write, and storage.
	SearchList []string `json:"search_list" required:"true"`
}

type TrafficStatistic struct {
	IndexTraffic       float64 `json:"index_traffic"`
	Storage            float64 `json:"storage"`
	WriteTraffic       float64 `json:"write_traffic"`
	LogGroupID         string  `json:"log_group_id"`
	LogGroupName       string  `json:"log_group_name"`
	LogStreamID        string  `json:"log_stream_id"`
	LogStreamName      string  `json:"log_stream_name"`
	LogGroupNameAlias  string  `json:"log_group_name_alias"`
	LogStreamNameAlias string  `json:"log_stream_name_alias"`
}

func ListTopNTrafficStatistics(
	client *golangsdk.ServiceClient,
	opts ListTopNTrafficStatisticsOpts,
) ([]TrafficStatistic, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/lts/topn-traffic-statistics
	raw, err := client.Post(client.ServiceURL("lts", "topn-traffic-statistics"), b, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type": "application/json;charset=UTF-8",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []TrafficStatistic
	err = extract.IntoSlicePtr(raw.Body, &res, "results")
	return res, err
}
