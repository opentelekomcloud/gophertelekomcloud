package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateOpts provides options used to create a user.
type UpdateOpts struct {
	// Name is the name of the new user.
	Name string `json:"name,omitempty"`

	// Password is the password of the new user.
	Password string `json:"password,omitempty"`

	// Email address with a maximum of 255 characters
	Email string `json:"email,omitempty"`

	// AreaCode is a country code, must be used together with Phone.
	AreaCode string `json:"areacode,omitempty"`

	// Phone is a mobile number with a maximum of 32 digits, must be used together with AreaCode.
	Phone string `json:"phone,omitempty"`

	// Description is a description of the user.
	Description string `json:"description,omitempty"`

	// AccessMode is the access type for IAM user
	AccessMode string `json:"access_mode,omitempty"`

	// Enabled sets the user status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// PasswordReset Indicates whether password reset is required at the first login.
	// By default, password reset is true.
	PasswordReset *bool `json:"pwd_status,omitempty"`
}

func ModifyUser(client *golangsdk.ServiceClient, id string, opts UpdateOpts) (*User, error) {
	b, err := build.RequestBody(opts, "user")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-USER/users/{user_id}
	raw, err := client.Put(client.ServiceURL("OS-USER", "users", id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res User
	return &res, extract.IntoStructPtr(raw.Body, &res, "user")
}
