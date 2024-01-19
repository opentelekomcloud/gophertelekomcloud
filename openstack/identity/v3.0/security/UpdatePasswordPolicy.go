package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePasswordPolicyOpts struct {
	// Maximum number of times that a character is allowed to consecutively present in a password.
	MaximumConsecutiveIdenticalChars int `json:"maximum_consecutive_identical_chars,omitempty"`
	// Maximum number of characters that a password can contain.
	MaximumPasswordLength int `json:"maximum_password_length,omitempty"`
	// Minimum period (minutes) after which users are allowed to make a password change.
	MinimumPasswordAge int `json:"minimum_password_age,omitempty"`
	// Minimum number of characters that a password must contain.
	MinimumPasswordLength int `json:"minimum_password_length,omitempty"`
	// Number of previously used passwords that are not allowed.
	NumberOfRecentPasswordsDisallowed int `json:"number_of_recent_passwords_disallowed,omitempty"`
	// Indicates whether the password can be the username or the username spelled backwards.
	PasswordNotUsernameOrInvert *bool `json:"password_not_username_or_invert,omitempty"`
	// Characters that a password must contain.
	PasswordRequirements string `json:"password_requirements,omitempty"`
	// Password validity period (days).
	PasswordValidityPeriod int `json:"password_validity_period,omitempty"`
}

func UpdatePasswordPolicy(client *golangsdk.ServiceClient, id string, opts UpdatePasswordPolicyOpts) (*PasswordPolicy, error) {
	b, err := build.RequestBody(opts, "password_policy")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/password-policy
	raw, err := client.Put(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "password-policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res PasswordPolicy
	return &res, extract.IntoStructPtr(raw.Body, &res, "password_policy")
}
