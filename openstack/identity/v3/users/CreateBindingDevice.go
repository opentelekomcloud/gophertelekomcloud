package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

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

func CreateBindingDevice(client *golangsdk.ServiceClient, opts BindMfaDevice) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// PUT /v3.0/OS-MFA/mfa-devices/bind
	_, err = client.Put(v30(client.ServiceURL("OS-MFA", "virtual-mfa-devices", "bind")), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
