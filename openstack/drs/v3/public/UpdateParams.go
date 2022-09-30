package public

type UpdateParamsRequest struct {
	// Task ID.
	JobId string
	// Parameter Groups Values: common performance
	Group string `json:"group"`
	// Information about the parameters to be modified.
	Params []ParamsReqBean `json:"params"`
}

type ParamsReqBean struct {
	// Database parameter name.
	Key string `json:"key"`
	// Parameter value of the destination database.
	TargetValue string `json:"target_value"`
}

// POST /v3/{project_id}/jobs/{job_id}/params

type UpdateParamsResponse struct {
	// Whether the parameters are modified.
	Success bool `json:"success,omitempty"`
	// Whether the instance needs to be restarted.
	// Values: true false
	ShouldRestart string `json:"should_restart,omitempty"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMsg string `json:"error_msg,omitempty"`
}
