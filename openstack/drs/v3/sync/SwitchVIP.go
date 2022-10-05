package sync

type SwitchVIPOpts struct {
	// Switchover type. Values:
	// source: Switch over the virtual IP address of the source database.
	// target: Switch over the virtual IP address of the destination database.
	Mode string `json:"mode"`
}

// POST  /v3/{project_id}/jobs/{job_id}/switch-vip

// IdJobResp
