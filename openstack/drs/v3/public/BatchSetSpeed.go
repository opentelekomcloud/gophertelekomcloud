package public

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

// PUT /v3/{project_id}/jobs/batch-limit-speed

// BatchJobsResponse
