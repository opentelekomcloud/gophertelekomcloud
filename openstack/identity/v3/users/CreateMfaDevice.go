package users

type CreateMfaDeviceReq struct {
	VirtualMfaDevice CreateMfaDevice `json:"virtual_mfa_device"`
}

type CreateMfaDevice struct {
	// Device name.
	// Minimum length: 1 character
	// Maximum length: 64 characters
	Name string `json:"name"`
	// ID of the user for whom you will create the MFA device.
	UserId string `json:"user_id"`
}

// POST /v3.0/OS-MFA/virtual-mfa-devices

type CreateMfaDeviceResponse struct {
	VirtualMfaDevice CreateMfaDeviceRespon `json:"virtual_mfa_device,omitempty"`
}

type CreateMfaDeviceRespon struct {
	// Serial number of the MFA device.
	SerialNumber string `json:"serial_number"`
	// Base32 seed, which a third-party system can use to generate a CAPTCHA code.
	Base32StringSeed string `json:"base32_string_seed"`
}
