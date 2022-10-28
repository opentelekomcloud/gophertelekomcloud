package users

type BindMfaDevice struct {
	// ID of the user to whom you will bind the virtual MFA device.
	UserId string `json:"user_id"`
	// Serial number of the virtual MFA device.
	SerialNumber string `json:"serial_number"`
	// Verification code 1.
	AuthenticationCodeFirst string `json:"authentication_code_first"`
	// Verification code 2.
	AuthenticationCodeSecond string `json:"authentication_code_second"`
}

// PUT /v3.0/OS-MFA/mfa-devices/bind
