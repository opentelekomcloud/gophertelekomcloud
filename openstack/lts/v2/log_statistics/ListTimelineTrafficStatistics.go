package log_statistics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListTimelineTrafficStatisticsOpts struct {
	// Timezone is the time zone used to aggregate the results.
	Timezone string `json:"-" q:"timezone" required:"true"`
	// StartTime is the start timestamp in milliseconds.
	StartTime *int64 `json:"start_time" required:"true"`
	// EndTime is the end timestamp in milliseconds.
	EndTime *int64 `json:"end_time" required:"true"`
	// Period is the query interval in hours, from 1 to 24.
	Period int `json:"period" required:"true"`
	// ResourceType is log_group, log_stream, or tenant.
	ResourceType string `json:"resource_type" required:"true"`
	// SearchType is write, index, or storage.
	SearchType string `json:"search_type" required:"true"`
	// ResourceID is the optional log group or log stream ID.
	ResourceID string `json:"resource_id,omitempty"`
}

type TimelineStatistic struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

func ListTimelineTrafficStatistics(
	client *golangsdk.ServiceClient,
	opts ListTimelineTrafficStatisticsOpts,
) ([]TimelineStatistic, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url, err := golangsdk.NewURLBuilder().
		WithEndpoints("lts", "timeline-traffic-statistics").
		WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/lts/timeline-traffic-statistics
	raw, err := client.Post(client.ServiceURL(url.String()), b, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type": "application/json;charset=UTF-8",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res []TimelineStatistic
	err = extract.IntoSlicePtr(raw.Body, &res, "results")
	return res, err
}
