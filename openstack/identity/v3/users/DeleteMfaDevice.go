package users

type DeleteMfaDeviceRequest struct {
	// ID of the user whose virtual MFA device is to be deleted, that is, the administrator's user ID.
	UserId string `json:"user_id"`
	// Serial number of the virtual MFA device.
	SerialNumber string `json:"serial_number"`
}

// DELETE /v3.0/OS-MFA/virtual-mfa-devices
