package v1

type ResetPasswordRequest struct {
	//
	ClusterId string                   `json:"cluster_id"`
	Body      ResetPasswordRequestBody `json:"body,omitempty"`
}

type ResetPasswordRequestBody struct {
	//
	NewPassword string `json:"new_password"`
}

// POST /v1.0/{project_id}/clusters/{cluster_id}/reset-password

type ResetPasswordResponse struct {
}
