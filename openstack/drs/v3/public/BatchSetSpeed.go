package public

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type BatchLimitSpeedReq struct {
	SpeedLimits []LimitSpeedReq `json:"speed_limits"`
}

type LimitSpeedReq struct {
	// Task ID.
	JobId string `json:"job_id"`
	// Request body of flow control information.
	SpeedLimit []SpeedLimitInfo `json:"speed_limit"`
}

type SpeedLimitInfo struct {
	// Start time (UTC) of flow control. The start time is an integer in hh:mm format and the minutes part is ignored.
	// hh indicates the hour, for example, 01:00.
	Begin string `json:"begin"`
	// End time (UTC) in the format of hh:mm, for example, 15:59. The value must end with 59.
	End string `json:"end"`
	// Speed. The value ranges from 1 to 9,999, in MB/s.
	Speed string `json:"speed"`
}

func BatchSetSpeed(client *golangsdk.ServiceClient, opts BatchUpdateDatabaseObjectReq) (*BatchJobsResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /v3/{project_id}/jobs/batch-limit-speed
	raw, err := client.Put(client.ServiceURL("jobs", "batch-limit-speed"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res BatchJobsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
