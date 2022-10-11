package cluster

type ResetPasswordRequest struct {
	// ID of the cluster whose password is to be reset. For details about how to obtain the ID, see 7.6 Obtaining the Cluster ID.
	ClusterId string `json:"cluster_id"`

	Body ResetPasswordOpts `json:"body,omitempty"`
}

type ResetPasswordOpts struct {
	// New password of the GaussDB(DWS) cluster administrator
	// A password must conform to the following rules:
	// Contains 8 to 32 characters.
	// Cannot be the same as the username or the username written in reverse order.
	// Contains at least three types of the following:
	// Lowercase letters
	// Uppercase letters
	// Digits
	// Special characters: ~!?,.:;-_'"(){}[]/<>@#%^&*+|\=
	// Cannot be the same as previous passwords.
	// Cannot be a weak password.
	NewPassword string `json:"new_password"`
}

// POST /v1.0/{project_id}/clusters/{cluster_id}/reset-password

type ResetPasswordResponse struct {
}
