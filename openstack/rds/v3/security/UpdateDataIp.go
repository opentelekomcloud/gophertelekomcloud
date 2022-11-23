package security

type UpdateDataIpOpts struct {
	InstanceId string
	// Indicates the private IP address.
	NewIp string `json:"new_ip"`
}

// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/ip

// workflowId 200
