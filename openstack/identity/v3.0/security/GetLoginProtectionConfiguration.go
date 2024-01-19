package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetLoginProtectionConfiguration(client *golangsdk.ServiceClient, id string) (*LoginProtectionConfig, error) {
	// GET /v3.0/OS-USER/users/{user_id}/login-protect
	raw, err := client.Get(client.ServiceURL("OS-USER", "users", id, "login-protect"), nil, nil)
	if err != nil {
		return nil, err
	}
	var res LoginProtectionConfig
	err = extract.IntoStructPtr(raw.Body, &res, "login_protect")
	return &res, err
}

type LoginProtectionConfig struct {
	// Indicates whether login protection has been enabled for a user. The value can be true or false.
	Enabled *bool `json:"enabled"`
	// User ID.
	UserId string `json:"user_id"`
	// Login authentication method of the user.
	VerificationMethod string `json:"verification_method"`
}
