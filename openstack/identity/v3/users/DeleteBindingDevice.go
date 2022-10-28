package users

type UnbindMfaDevice struct {
	// ID of the user from whom you will unbind the MFA device.
	UserId string `json:"user_id"`
	// Administrator: Set this parameter to any value, because verification is not required.
	// IAM user: Enter the MFA verification code.
	AuthenticationCode string `json:"authentication_code"`
	// Serial number of the MFA device.
	SerialNumber string `json:"serial_number"`
}

// PUT /v3.0/OS-MFA/mfa-devices/unbind
