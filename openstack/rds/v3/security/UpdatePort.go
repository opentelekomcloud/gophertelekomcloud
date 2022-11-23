package security

type UpdatePortOpts struct {
	InstanceId string
	// Specifies port information for all DB engines.
	// The MySQL database port ranges from 1024 to 65535 (excluding 12017 and 33071, which are occupied by the RDS system and cannot be used).
	// The PostgreSQL database port ranges from 2100 to 9500.
	// The Microsoft SQL Server database port is 1433 or ranges from 2100 to 9500 (excluding 5355 and 5985).
	// The default values is as follows:
	// The default value of MySQL is 3306.
	// The default value of PostgreSQL is 5432.
	// The default value of Microsoft SQL Server is 1433.
	Port int32 `json:"port"`
}

// PUT https://{Endpoint}/v3/{project_id}/instances/{instance_id}/port

// workflowId 200
