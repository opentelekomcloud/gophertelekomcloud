package public

type BatchQueryParamReq struct {
	// Request body for querying tasks in batches.
	Jobs []string `json:"jobs"`
	// Whether to obtain database parameters again. 1 indicates yes, and 0 indicates no
	// (obtaining parameters from the cache).
	// Set this parameter to 1 when this API is called for the first time.
	Refresh string `json:"refresh"`
}

// POST /v3/{project_id}/jobs/batch-get-params

type BatchShowParamsResponse struct {
	ParamsList []QueryDbParamsResp `json:"params_list,omitempty"`
	Count      int32               `json:"count,omitempty"`
}

type QueryDbParamsResp struct {
	Params []Params `json:"params,omitempty"`
}

type Params struct {
	// Parameter comparison result. Values: true false
	CompareResult string `json:"compare_result,omitempty"`
	// Type
	DataType string `json:"data_type,omitempty"`
	// Metric Type Values:
	// common: common parameter.
	// performance: performance parameter.
	Group string `json:"group,omitempty"`
	// Parameter name
	Key string `json:"key,omitempty"`
	// Whether the instance needs to be restarted.
	// Values: true false
	NeedRestart string `json:"need_restart,omitempty"`
	// Source database parameter value.
	SourceValue string `json:"source_value,omitempty"`
	// Parameter value of the destination database.
	TargetValue string `json:"target_value,omitempty"`
	// Value Range
	ValueRange string `json:"value_range,omitempty"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMessage string `json:"error_message,omitempty"`
}
