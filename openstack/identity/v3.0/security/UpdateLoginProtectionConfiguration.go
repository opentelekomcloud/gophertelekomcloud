package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type LoginProtectionUpdateOpts struct {
	// Indicates whether login protection has been enabled for the user. The value can be true or false.
	Enabled *bool `json:"enabled" required:"true"`
	// Login authentication method of the user. Options: sms, email, and vmfa.
	VerificationMethod string `json:"verification_method" required:"true"`
}

func UpdateLoginProtectionConfiguration(client *golangsdk.ServiceClient, id string, opts LoginProtectionUpdateOpts) (*LoginProtectionConfig, error) {
	b, err := build.RequestBody(opts, "login_protect")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-USER/users/{user_id}/login-protect
	raw, err := client.Put(client.ServiceURL("OS-USER", "users", id, "login-protect"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res LoginProtectionConfig
	return &res, extract.IntoStructPtr(raw.Body, &res, "login_protect")
}
