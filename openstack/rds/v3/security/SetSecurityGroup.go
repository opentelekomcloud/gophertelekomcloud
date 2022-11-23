package security

type SetSecurityGroupOpts struct {
	InstanceId string
	// Specifies the security group ID.
	SecurityGroupId string `json:"security_group_id"`
}

// PUT /v3/{project_id}/instances/{instance_id}/security-group

// WorkflowId 200
