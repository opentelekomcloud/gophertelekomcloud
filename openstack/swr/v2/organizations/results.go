package organizations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type CreateResult struct {
	golangsdk.ErrResult
}

type DeleteResult struct {
	golangsdk.ErrResult
}

type ListResult struct {
	golangsdk.Result
}

type GetResult struct {
	golangsdk.Result
}

type Organization struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CreatorName string `json:"creator_name"`
	// Auth - user permission. The value can be `7`: manage, `3`: write, `1`: read
	Auth int `json:"auth"`
}

type OrganizationPage struct {
	pagination.SinglePageBase
}

func ExtractOrganizations(p pagination.Page) ([]Organization, error) {
	var orgs []Organization
	err := p.(OrganizationPage).ExtractIntoSlicePtr(&orgs, "namespaces")
	return orgs, err
}

func (r GetResult) Extract() (*Organization, error) {
	org := new(Organization)
	err := r.ExtractIntoStructPtr(org, "")
	return org, err
}
