package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type DeleteMfaDeviceOpts struct {
	// ID of the user whose virtual MFA device is to be deleted, that is, the administrator's user ID.
	UserId string `json:"user_id"`
	// Serial number of the virtual MFA device.
	SerialNumber string `json:"serial_number"`
}

func DeleteMfaDevice(client *golangsdk.ServiceClient, opts DeleteMfaDeviceOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// DELETE /v3.0/OS-MFA/virtual-mfa-devices
	_, err = client.DeleteWithBody(v30(client.ServiceURL("OS-MFA", "virtual-mfa-devices")), b, nil)
	return
}
