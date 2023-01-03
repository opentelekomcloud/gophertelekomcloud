package organizations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

type GetResult struct {
	golangsdk.Result
}

type CreatePermissionsResult struct {
	golangsdk.ErrResult
}

type DeletePermissionsResult struct {
	golangsdk.ErrResult
}

type UpdatePermissionsResult struct {
	golangsdk.ErrResult
}

type Auth struct {
	// User ID, which needs to be obtained from the IAM service.
	UserID string `json:"user_id"`
	// Username, which needs to be obtained from the IAM service.
	Username string `json:"user_name"`
	// User permission that is configured. The value can be 1, 3, or 7. 7: manage 3: write 1: read
	Auth int `json:"auth"`
}

type OrganizationPermissions struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CreatorName string `json:"creator_name"`
	SelfAuth    Auth   `json:"self_auth"`
	OthersAuth  []Auth `json:"others_auths"`
}

type GetPermissionsResult struct {
	golangsdk.Result
}

func (r GetPermissionsResult) Extract() (*OrganizationPermissions, error) {
	perm := new(OrganizationPermissions)
	err := r.ExtractIntoStructPtr(perm, "")
	if err != nil {
		return nil, err
	}
	return perm, nil
}
