package users

// GET /v3.0/OS-MFA/virtual-mfa-devices

type ListUserMfaDevicesResponse struct {
	// 虚拟MFA设备信息列表。
	VirtualMfaDevices []MfaDeviceResult `json:"virtual_mfa_devices,omitempty"`
}

type MfaDeviceResult struct {
	// Virtual MFA device serial number.
	SerialNumber string `json:"serial_number"`
	// User ID.
	UserId string `json:"user_id"`
}
