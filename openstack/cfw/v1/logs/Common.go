package logs

type LogConfigOpts struct {
	// Firewall ID, which can be obtained by referring to Obtaining a Firewall ID.
	FWInstanceID string `json:"fw_instance_id" required:"true"`

	// Whether to enable LTS: 1 (yes), 0 (no).
	// If set to 1, then LtsAttackLogStreamEnable, LtsAccessLogStreamEnable,
	// and LtsFlowLogStreamEnable must be mandatory.
	LtsEnable int `json:"lts_enable" required:"true"`

	// Log Tank Service (LTS) log group ID, which can be obtained by calling the
	// API for querying all the log groups of an account in LTS.
	// Find the value in log_groups.log_group_id (the period [.] is used to
	// separate different levels of objects).
	LtsLogGroupID string `json:"lts_log_group_id" required:"true"`

	// Attack log stream ID, which can be obtained by calling the API for
	// querying all the log streams in a specified log group in LTS.
	// Find the value in log_streams.log_stream_id (the period [.] is used to
	// separate different levels of objects).
	LtsAttackLogStreamID string `json:"lts_attack_log_stream_id,omitempty"`

	// Whether to enable the attack log stream: 1 (yes), 0 (no).
	LtsAttackLogStreamEnable int `json:"lts_attack_log_stream_enable"`

	// Access control log stream ID, which can be obtained by calling the API
	// for querying all the log streams in a specified log group in LTS.
	// Find the value in log_streams.log_stream_id (the period [.] is used to
	// separate different levels of objects).
	LtsAccessLogStreamID string `json:"lts_access_log_stream_id,omitempty"`

	// Whether to enable the access control stream: 1 (yes), 0 (no).
	LtsAccessLogStreamEnable int `json:"lts_access_log_stream_enable"`

	// Traffic log ID, which can be obtained by calling the API for querying
	// all the log streams in a specified log group in LTS.
	// Find the value in log_streams.log_stream_id (the period [.] is used to
	// separate different levels of objects).
	LtsFlowLogStreamID string `json:"lts_flow_log_stream_id,omitempty"`

	// Whether to enable the traffic log function: 1 (yes), 0 (no).
	LtsFlowLogStreamEnable int `json:"lts_flow_log_stream_enable"`
}

/*
########################## RESPONSE STRUCTS ##############################
*/

type DataResponse struct {
	// Return value for updating log configurations. The value is the firewall ID.
	Data string `json:"data"`
}
