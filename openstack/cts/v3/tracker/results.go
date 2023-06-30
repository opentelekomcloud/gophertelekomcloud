package tracker

import (
	"net/http"

	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type Tracker struct {
	// Unique tracker ID.
	Id string `json:"id"`
	// Timestamp when the tracker was created.
	CreateTime int64 `json:"create_time"`
	// Trace analysis
	Lts Lts `json:"lts"`
	// Tracker type
	TrackerType string `json:"tracker_type"`
	// Account ID
	DomainId string `json:"domain_id"`
	// Project id
	ProjectId string `json:"project_id"`
	// Tracker name
	TrackerName string `json:"tracker_name"`
	// Tracker status
	Status string `json:"status"`
	// This parameter is returned only when the tracker status is error.
	Detail string `json:"detail"`
	// Tracker type
	ObsInfo ObsInfo `json:"obs_info"`
}

type Lts struct {
	// Whether trace analysis is enabled.
	IsLtsEnabled bool `json:"is_lts_enabled"`
	// Name of the Log Tank Service (LTS) log group.
	LogGroupName string `json:"log_group_name"`
	// Name of the LTS log stream.
	LogTopicName string `json:"log_topic_name"`
}

func extra(err error, raw *http.Response) (*Tracker, error) {
	if err != nil {
		return nil, err
	}

	var res Tracker
	err = extract.Into(raw.Body, &res)
	return &res, err
}

func extraStruct(err error, raw *http.Response) ([]Tracker, error) {
	if err != nil {
		return nil, err
	}
	var res []Tracker

	err = extract.IntoSlicePtr(raw.Body, &res, "trackers")
	return res, err
}
