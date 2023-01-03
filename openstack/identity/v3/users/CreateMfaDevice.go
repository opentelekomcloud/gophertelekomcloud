package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateMfaDeviceOpts struct {
	// Device name.
	// Minimum length: 1 character
	// Maximum length: 64 characters
	Name string `json:"name"`
	// ID of the user for whom you will create the MFA device.
	UserId string `json:"user_id"`
}

func CreateMfaDevice(client *golangsdk.ServiceClient, opts CreateMfaDeviceOpts) (*CreateMfaDeviceResponse, error) {
	b, err := build.RequestBody(opts, "virtual_mfa_device")
	if err != nil {
		return nil, err
	}

	// POST /v3.0/OS-MFA/virtual-mfa-devices
	raw, err := client.Post(v30(client.ServiceURL("OS-MFA", "virtual-mfa-devices")), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res CreateMfaDeviceResponse
	return &res, extract.IntoStructPtr(raw.Body, &res, "virtual_mfa_device")
}

type CreateMfaDeviceResponse struct {
	// Serial number of the MFA device.
	SerialNumber string `json:"serial_number"`
	// Base32 seed, which a third-party system can use to generate a CAPTCHA code.
	Base32StringSeed string `json:"base32_string_seed"`
}
