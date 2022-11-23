package security

type SwitchSslOpts struct {
	// Specifies the DB instance ID.
	InstanceId string
	// Specifies whether to enable SSL.
	// true: SSL is enabled.
	// false: SSL is disabled.
	SslOption bool `json:"ssl_option"`
}

// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/ssl

// 200
