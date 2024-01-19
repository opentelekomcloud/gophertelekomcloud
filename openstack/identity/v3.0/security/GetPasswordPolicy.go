package security

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetPasswordPolicy(client *golangsdk.ServiceClient, id string) (*PasswordPolicy, error) {
	// GET /v3.0/OS-SECURITYPOLICY/domains/{domain_id}/password-policy
	raw, err := client.Get(client.ServiceURL("OS-SECURITYPOLICY", "domains", id, "password-policy"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res PasswordPolicy
	err = extract.IntoStructPtr(raw.Body, &res, "password_policy")
	return &res, err
}

type PasswordPolicy struct {
	// Maximum number of times that a character is allowed to consecutively present in a password.
	MaximumConsecutiveIdenticalChars int `json:"maximum_consecutive_identical_chars"`
	// Maximum number of characters that a password can contain.
	MaximumPasswordLength int `json:"maximum_password_length"`
	// Minimum period (minutes) after which users are allowed to make a password change.
	MinimumPasswordAge int `json:"minimum_password_age"`
	// Minimum number of characters that a password must contain.
	MinimumPasswordLength int `json:"minimum_password_length"`
	// Number of previously used passwords that are not allowed.
	NumberOfRecentPasswordsDisallowed int `json:"number_of_recent_passwords_disallowed"`
	// Indicates whether the password can be the username or the username spelled backwards.
	PasswordNotUsernameOrInvert *bool `json:"password_not_username_or_invert"`
	// Characters that a password must contain.
	PasswordRequirements string `json:"password_requirements"`
	// Password validity period (days).
	PasswordValidityPeriod int `json:"password_validity_period"`
}
