package public

type BatchTestConnectionOpts struct {
	Jobs []TestEndPoint `json:"jobs"`
}

type TestEndPoint struct {
	// Task ID.
	Id string `json:"id"`
	// Network type.
	// Value: vpn vpc eip
	NetType string `json:"net_type"`
	// Database type.
	// Value: mysql mongodb taurus postgresql
	DbType string `json:"db_type"`
	// Database IP address.
	Ip string `json:"ip"`
	// Database port number. This parameter must be set to 0 for the Mongo and DDS databases.
	DbPort int32 `json:"db_port,omitempty"`
	// The RDS or GaussDB(for MySQL) instance ID. This parameter is mandatory for RDS or GaussDB(for MySQL) instances.
	InstId string `json:"inst_id,omitempty"`
	// Database account.
	DbUser string `json:"db_user"`
	// Database password.
	DbPassword string `json:"db_password"`
	// Whether SSL is enabled.
	SslLink bool `json:"ssl_link,omitempty"`
	// The SSL certificate content, which is encrypted using Base64.
	// This parameter is mandatory for secure connection to the source database.
	SslCertKey string `json:"ssl_cert_key,omitempty"`
	// The SSL certificate name. This parameter is mandatory for secure connection to the source database.
	SslCertName string `json:"ssl_cert_name,omitempty"`
	// The checksum value of the SSL certificate, which is used for backend verification.
	// This parameter is mandatory for secure connection to the source database.
	SslCertCheckSum string `json:"ssl_cert_check_sum,omitempty"`
	// The SSL certificate password. The certificate file name extension is .p12 and requires a password.
	SslCertPassword string `json:"ssl_cert_password,omitempty"`
	// ID of the VPC where the instance is located.
	VpcId string `json:"vpc_id,omitempty"`
	// ID of the subnet where the instance is located.
	SubnetId string `json:"subnet_id,omitempty"`
	// Source database: so. Destination database: ta. Default value: so
	// Values: so ta
	EndPointType string `json:"end_point_type"`
	// Region of the RDS DB instance. This parameter is mandatory when the RDS DB instance is used.
	Region string `json:"region,omitempty"`
	// Project ID of the region where the user is located. This parameter is mandatory when the RDS DB instance is used.
	ProjectId string `json:"project_id,omitempty"`
	// Database username, which is the DDS authentication database or the service name of the Oracle database.
	DbName string `json:"db_name,omitempty"`
}

// POST /v3/{project_id}/jobs/batch-connection

type BatchValidateConnectionsResponse struct {
	Results []CheckJobResp `json:"results,omitempty"`
	Count   int32          `json:"count,omitempty"`
}

type CheckJobResp struct {
	// Task ID.
	Id string `json:"id"`
	// Test result. Value:
	// success: indicates that the connection test is successful.
	// failed: indicates that the connection test fails.
	Status string `json:"status"`
	// Error code, which is optional and indicates the returned information about the failure status.
	ErrorCode string `json:"error_code,omitempty"`
	// Error message, which is optional and indicates the returned information about the failure status.
	ErrorMsg string `json:"error_msg,omitempty"`
	// Whether the request is successful.
	Success bool `json:"success,omitempty"`
}
