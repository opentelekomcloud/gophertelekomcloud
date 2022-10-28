package users

type ShowUserMfaDeviceRequest struct {
	UserId string `json:"user_id"`
}

// GET /v3.0/OS-MFA/users/{user_id}/virtual-mfa-device

type ShowUserMfaDeviceResponse struct {
	VirtualMfaDevice MfaDeviceResult `json:"virtual_mfa_device,omitempty"`
}
