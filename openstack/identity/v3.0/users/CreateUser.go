package users

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts provides options used to create a user.
type CreateOpts struct {
	// Name is the name of the new user.
	Name string `json:"name" required:"true"`

	// DomainID is the ID of the domain the user belongs to.
	DomainID string `json:"domain_id" required:"true"`

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

func CreateUser(client *golangsdk.ServiceClient, opts CreateOpts) (*User, error) {
	b, err := build.RequestBody(opts, "user")
	if err != nil {
		return nil, err
	}

	// POST /v3.0/OS-USER/users
	raw, err := client.Post(client.ServiceURL("OS-USER", "users"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res User
	return &res, extract.IntoStructPtr(raw.Body, &res, "user")
}

// User represents a User in the OpenStack Identity Service.
type User struct {
	ID                string `json:"id"`
	DomainID          string `json:"domain_id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	AreaCode          string `json:"areacode"`
	Phone             string `json:"phone"`
	Description       string `json:"description"`
	AccessMode        string `json:"access_mode"`
	Enabled           bool   `json:"enabled"`
	PasswordStatus    bool   `json:"pwd_status"`
	PasswordStrength  string `json:"pwd_strength"`
	PasswordExpiresAt string `json:"password_expires_at"`
	IsDomainOwner     bool   `json:"is_domain_owner"`
	CreateAt          string `json:"create_time"`
	UpdateAt          string `json:"update_time"`
	LastLogin         string `json:"last_login_time"`
	Status            string `json:"status"`
	XuserID           string `json:"xuser_id"`
	XuserType         string `json:"xuser_type"`
}
