package log_statistics

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListLogHistogramOpts struct {
	// StartTime is the start timestamp accurate to a millisecond.
	StartTime string `json:"start_time" required:"true"`
	// EndTime is the end timestamp accurate to a millisecond.
	EndTime string `json:"end_time" required:"true"`
	// StepInterval is the time step in milliseconds.
	StepInterval int64 `json:"step_interval" required:"true"`
	// GroupID is the log group ID.
	GroupID string `json:"group_id" required:"true"`
	// StreamID is the log stream ID.
	StreamID string `json:"stream_id" required:"true"`
	// Keyword is the keyword to count.
	Keyword string `json:"key_word" required:"true"`
	// IsIterative indicates whether the log query is iterative.
	IsIterative *bool `json:"is_iterative,omitempty"`
}

type LogHistogramResponse struct {
	Count           int64          `json:"count"`
	Histogram       []LogHistogram `json:"histogram"`
	IsQueryComplete bool           `json:"isQueryComplete"`
}

type LogHistogram struct {
	Num       int64 `json:"num"`
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
}

func ListLogHistogram(client *golangsdk.ServiceClient, opts ListLogHistogramOpts) (*LogHistogramResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /v2/{project_id}/lts/keyword-count
	raw, err := client.Post(client.ServiceURL("lts", "keyword-count"), b, nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{
			"Content-Type": "application/json;charset=UTF-8",
		},
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res LogHistogramResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
