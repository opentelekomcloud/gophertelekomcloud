package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowUserMfaDevice(client *golangsdk.ServiceClient, userId string) (*MfaDeviceResult, error) {
	// GET /v3.0/OS-MFA/users/{user_id}/virtual-mfa-device
	raw, err := client.Get(v30(client.ServiceURL("OS-MFA", "users", userId, "virtual-mfa-device")), nil, nil)
	if err != nil {
		return nil, err
	}

	var res MfaDeviceResult
	err = extract.IntoStructPtr(raw.Body, &res, "virtual_mfa_device")
	return &res, err
}
