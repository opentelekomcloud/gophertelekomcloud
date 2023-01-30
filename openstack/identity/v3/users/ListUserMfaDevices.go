package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListUserMfaDevices(client *golangsdk.ServiceClient) ([]MfaDeviceResult, error) {
	// GET /v3.0/OS-MFA/virtual-mfa-devices
	raw, err := client.Get(v30(client.ServiceURL("OS-MFA", "virtual-mfa-devices")), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []MfaDeviceResult
	err = extract.IntoSlicePtr(raw.Body, &res, "virtual_mfa_devices")
	return res, err
}

type MfaDeviceResult struct {
	// Virtual MFA device serial number.
	SerialNumber string `json:"serial_number"`
	// User ID.
	UserId string `json:"user_id"`
}
