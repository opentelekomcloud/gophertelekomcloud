package organizations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
)

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
