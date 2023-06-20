package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
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
	// Duration that traces are stored in the OBS bucket. When tracker_type is set to system,
	// the default value is 0, indicating permanent storage.
	BucketLifecycle *int `json:"bucket_lifecycle,omitempty"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Tracker, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /3/{project_id}/tracker
	raw, err := client.Post(client.ServiceURL("tracker"), b, nil, nil)
	return extra(err, raw)
}
