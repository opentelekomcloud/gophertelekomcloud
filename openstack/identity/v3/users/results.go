package users

import (
	"encoding/json"
	"time"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

// User represents a User in the OpenStack Identity Service.
type User struct {
	// DefaultProjectID is the ID of the default project of the user.
	DefaultProjectID string `json:"default_project_id"`

	// Description is a description of the user.
	Description string `json:"description"`

	// DomainID is the domain ID the user belongs to.
	DomainID string `json:"domain_id"`

	// Enabled is whether the user is enabled.
	Enabled bool `json:"enabled"`

	// ID is the unique ID of the user.
	ID string `json:"id"`

	// Links contains referencing links to the user.
	Links map[string]interface{} `json:"links"`

	// Name is the name of the user.
	Name string `json:"name"`

	// PasswordExpiresAt is the timestamp when the user's password expires.
	PasswordExpiresAt time.Time `json:"-"`

	// Email is the email of the user
	Email string `json:"email,omitempty"`

	// AreaCode is country code
	AreaCode string `json:"areacode,omitempty"`

	// Phone is mobile number, which can contain a maximum of 32 digits.
	// The mobile number must be used together with a country code.
	Phone string `json:"phone,omitempty"`

	// Whether password reset is required at first login
	PwdResetRequired bool `json:"pwd_status,omitempty"`

	// XUserType is Type of the IAM user in the external system.
	XUserType string `json:"xuser_type,omitempty"`

	// XUserID is ID of the IAM user in the external system.
	XUserID string `json:"xuser_id,omitempty"`
}

func (r *User) UnmarshalJSON(b []byte) error {
	type tmp User
	var s struct {
		tmp
		PasswordExpiresAt golangsdk.JSONRFC3339MilliNoZ `json:"password_expires_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = User(s.tmp)

	r.PasswordExpiresAt = time.Time(s.PasswordExpiresAt)

	return err
}

type userResult struct {
	golangsdk.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a User.
type GetResult struct {
	userResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a User.
type CreateResult struct {
	userResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a User.
type UpdateResult struct {
	userResult
}

// UpdateExtendedResult is the response from an UpdateExtended operation. Call its Extract
// method to interpret it as a User.
type UpdateExtendedResult struct {
	userResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

// UserPage is a single page of User results.
type UserPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether a UserPage contains any results.
func (r UserPage) IsEmpty() (bool, error) {
	users, err := ExtractUsers(r)
	return len(users) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r UserPage) NextPageURL() (string, error) {
	var s struct {
		Links struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Links.Next, err
}

// ExtractUsers returns a slice of Users contained in a single page of results.
func ExtractUsers(r pagination.Page) ([]User, error) {
	var s []User
	err := (r.(UserPage)).ExtractIntoSlicePtr(&s, "users")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// Extract interprets any user results as a User.
func (r userResult) Extract() (*User, error) {
	s := new(User)
	err := r.ExtractIntoStructPtr(s, "user")
	if err != nil {
		return nil, err
	}
	return s, nil
}

type AddMembershipResult struct {
	golangsdk.ErrResult
}

type WelcomeResult struct {
	golangsdk.ErrResult
}
