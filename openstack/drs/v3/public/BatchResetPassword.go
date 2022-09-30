package public

type BatchModifyPwdReq struct {
	Jobs []ModifyPwdEndPoint `json:"jobs"`
}

type ModifyPwdEndPoint struct {
	// Database password.
	DbPassword string `json:"db_password"`
	// Type.
	// so indicates the source database.
	// ta indicates the destination database.
	// Values: so ta
	EndPointType string `json:"end_point_type"`
	// Task ID.
	JobId string `json:"job_id"`
	// Information required for Kerberos authentication.
	Kerberos KerberosVo `json:"kerberos,omitempty"`
}

type KerberosVo struct {
	// krb5 configuration file.
	Krb5ConfFile string `json:"krb5_conf_file,omitempty"`
	// Key file.
	KeyTabFile string `json:"key_tab_file,omitempty"`
	// Domain name.
	DomainName string `json:"domain_name,omitempty"`
	// Kerberos user object.
	UserPrincipal string `json:"user_principal,omitempty"`
}

// PUT /v3/{project_id}/jobs/batch-modify-pwd

type BatchResetPasswordResponse struct {
	Results []ModifyDbPwdResp `json:"results,omitempty"`
	Count   int32             `json:"count,omitempty"`
}

type ModifyDbPwdResp struct {
	Id           string `json:"id,omitempty"`
	Status       string `json:"status,omitempty"`
	EndPointType string `json:"end_point_type,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMsg     string `json:"error_msg,omitempty"`
}
