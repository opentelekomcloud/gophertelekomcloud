package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// ToInstanceCreateMap is used for type convert
func (ops CreateOps) ToInstanceCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(ops, "")
}

// ToInstanceUpdateMap is used for type convert
func (opts UpdateOpts) ToInstanceUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// UpdatePasswordOptsBuilder is an interface which can build the map paramter of update password function
type UpdatePasswordOptsBuilder interface {
	ToPasswordUpdateMap() (map[string]interface{}, error)
}

// UpdatePasswordOpts is a struct which represents the parameters of update function
type UpdatePasswordOpts struct {
	// Old password. It may be empty.
	OldPassword string `json:"old_password" required:"true"`
	// New password.
	// Password complexity requirements:
	// A string of 6–32 characters.
	// Must be different from the old password.
	// Contains at least two types of the following characters:
	// Uppercase letters
	// Lowercase letters
	// Digits
	// Special characters `~!@#$%^&*()-_=+\|[{}]:'",<.>/?
	NewPassword string `json:"new_password" required:"true"`
}

// ToPasswordUpdateMap is used for type convert
func (opts UpdatePasswordOpts) ToPasswordUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}
