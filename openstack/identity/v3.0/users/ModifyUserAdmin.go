package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// UpdateAdminOpts provides options used to update user as administrator
type UpdateAdminOpts struct {
	// Id is the IAM user ID
	Id string `json:"-"`

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

	// PwdStatus Indicates whether the user must change their password at the first login.
	PwdStatus *bool `json:"pwd_status,omitempty"`

	// XuserType is the type of the user in the external system.
	XuserType string `json:"xuser_type"`

	// XuserId is the ID of the user in the external system.
	XuserId string `json:"xuser_id"`
}

func ModifyUserAdmin(client *golangsdk.ServiceClient, opts UpdateAdminOpts) (*User, error) {
	b, err := build.RequestBody(opts, "user")
	if err != nil {
		return nil, err
	}

	// PUT /v3.0/OS-USER/users/{user_id}
	raw, err := client.Put(client.ServiceURL("OS-USER", "users", opts.Id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res User
	return &res, extract.IntoStructPtr(raw.Body, &res, "user")
}
