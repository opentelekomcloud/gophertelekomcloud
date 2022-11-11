package policies

type CreatePolicyRuleOpts struct {
	// Unique ID of the stream.
	StreamId string `json:"stream_id"`
	// Authorized users.
	// If the permission is granted to a specified tenant, the format is domainName.*.
	// If the permission is granted to a specified sub-user of a tenant, the format is domainName.userName.
	// Multiple accounts can be added and separated by commas (,),
	// for example, domainName1.userName1,do mainName2.userName2.
	PrincipalName string `json:"principal_name"`
	// Authorization operation type.
	// - putRecords: upload data.
	// - getRecords: download data.
	// Enumeration values:
	// putRecords
	// getRecords
	ActionType string `json:"action_type"`
	// Authorization impact type.
	// - accept: The authorization operation is allowed.
	// Enumeration values:
	// - accept
	Effect string `json:"effect"`
}

// POST /v2/{project_id}/streams/{stream_name}/policies

type CreatePolicyRuleResponse struct {
}
