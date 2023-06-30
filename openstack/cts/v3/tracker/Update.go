package tracker

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// Tracker type. The value can be system (management tracker).
	TrackerType string `json:"tracker_type" required:"true"`
	// Tracker name. When tracker_type is set to system, the default value system is used.
	TrackerName string `json:"tracker_name" required:"true"`
	// Configurations of an OBS bucket to which traces will be transferred.
	ObsInfo ObsInfo `json:"obs_info,omitempty"`
	// Indicates whether to enable trace analysis.
	IsLtsEnabled *bool `json:"is_lts_enabled,omitempty"`
	// Status of a tracker. The value can be enabled or disabled.
	// If you change the value to disabled, the tracker stops recording traces.
	Status string `json:"status,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*Tracker, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /3/{project_id}/tracker
	raw, err := client.Put(client.ServiceURL("tracker"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
