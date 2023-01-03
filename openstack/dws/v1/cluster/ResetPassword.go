package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ResetPasswordOpts struct {
	ClusterId string `json:"-"`
	// New password of the GaussDB(DWS) cluster administrator
	// A password must conform to the following rules:
	// Contains 8 to 32 characters.
	// Cannot be the same as the username or the username written in reverse order.
	// Contains at least three types of the following:
	// Lowercase letters
	// Uppercase letters
	// Digits
	// Special characters: ~!?,.:;-_'"(){}[]/<>@#%^&*+|\=
	// Cannot be the same as previous passwords.
	// Cannot be a weak password.
	NewPassword string `json:"new_password" required:"true"`
}

func ResetPassword(client *golangsdk.ServiceClient, opts ResetPasswordOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v1.0/{project_id}/clusters/{cluster_id}/reset-password
	_, err = client.Post(client.ServiceURL("clusters", opts.ClusterId, "reset-password"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
