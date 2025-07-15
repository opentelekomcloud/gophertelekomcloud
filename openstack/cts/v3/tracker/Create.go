package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Tracker type. The value can be system (management tracker).
	TrackerType string `json:"tracker_type" required:"true"`
	// Tracker name. When tracker_type is set to system, the default value system is used.
	TrackerName string `json:"tracker_name" required:"true"`
	// Configurations of an OBS bucket to which traces will be transferred.
	ObsInfo ObsInfo `json:"obs_info,omitempty"`
	// Indicates whether to enable trace analysis.
	IsLtsEnabled bool `json:"is_lts_enabled,omitempty"`
}

type ObsInfo struct {
	// OBS bucket name. The value contains 3 to 63 characters and must start with a digit or lowercase letter.
	// Only lowercase letters, digits, hyphens (-), and periods (.) are allowed.
	BucketName string `json:"bucket_name,omitempty"`
	// Prefix of trace files that need to be stored in OBS buckets.
	// The value can contain 0 to 64 characters, including letters, digits, hyphens (-), underscores (_), and periods (.).
	FilePrefixName string `json:"file_prefix_name,omitempty"`
	// Whether the OBS bucket is automatically created by the tracker.
	IsObsCreated *bool `json:"is_obs_created,omitempty"`
	// Compression type. The value can be JSON (no compression) or GZIP (compression). The default format is GZIP.
	CompressType string `json:"compress_type,omitempty"`
	// Whether to sort the path by cloud service. If this option is enabled, the cloud service name is added to the transfer file path.
	IsSortByService *bool `json:"is_sort_by_service,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Tracker, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /3/{project_id}/tracker
	raw, err := client.Post(client.ServiceURL("tracker"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res Tracker
	err = extract.Into(raw.Body, &res)
	return &res, err
}

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
