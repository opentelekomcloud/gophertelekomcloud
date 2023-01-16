package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UnbindMfaDevice struct {
	// ID of the user from whom you will unbind the MFA device.
	UserId string `json:"user_id"`
	// Administrator: Set this parameter to any value, because verification is not required.
	// IAM user: Enter the MFA verification code.
	AuthenticationCode string `json:"authentication_code"`
	// Serial number of the MFA device.
	SerialNumber string `json:"serial_number"`
}

func DeleteBindingDevice(client *golangsdk.ServiceClient, opts UnbindMfaDevice) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT /v3.0/OS-MFA/mfa-devices/unbind
	_, err = client.Put(v30(client.ServiceURL("OS-MFA", "virtual-mfa-devices", "unbind")), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
