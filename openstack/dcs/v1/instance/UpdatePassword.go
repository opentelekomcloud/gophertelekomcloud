package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdatePasswordOpts struct {
	OldPassword string `json:"old_password"`
	// New password.
	// Password complexity requirements:
	// Cannot be empty.
	// Cannot be the username or the username spelled backwards.
	// Can be 8 to 32 characters long.
	// Contain at least three of the following character types:
	// Lowercase letters
	// Uppercase letters
	// Digits
	// Special characters (`~!@#$^&*()-_=+\|{}:,<.>/?)
	NewPassword string `json:"new_password"`
}

func UpdatePassword(client *golangsdk.ServiceClient, instanceID string, opts UpdatePasswordOpts) (*UpdatePasswordResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("instances", instanceID, "password"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res UpdatePasswordResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type UpdatePasswordResponse struct {
	// Account lockout duration. If the old password is incorrect or the account is locked, the value of this parameter is not null.
	LockTime string `json:"lock_time,omitempty"`
	// An indicator of whether the password is successfully changed: Options:
	// Success: Password changed successfully.
	// passwordFailed: The old password is incorrect.
	// Locked: This account has been locked.
	// Failed: Failed to change the password.
	Result string `json:"result,omitempty"`
	// Remaining time before the account is unlocked. If the account is locked, the value of this parameter is not null.
	LockTimeLeft string `json:"lock_time_left,omitempty"`
	// Number of remaining password attempts. If the old password is incorrect, the value of this parameter is not null.
	RetryTimesLeft string `json:"retry_times_left,omitempty"`
	// Result of password change.
	Message string `json:"message,omitempty"`
}
