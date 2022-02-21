package agency

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/identity/v3/roles"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type Agency struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	DomainID            string `json:"domain_id"`
	DelegatedDomainID   string `json:"trust_domain_id"`
	DelegatedDomainName string `json:"trust_domain_name"`
	Description         string `json:"description"`
	Duration            string `json:"duration"`
	ExpireTime          string `json:"expire_time"`
	CreateTime          string `json:"create_time"`
}

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*Agency, error) {
	var s struct {
		Agency *Agency `json:"agency"`
	}
	err := r.ExtractInto(&s)
	return s.Agency, err
}

type GetResult struct {
	commonResult
}

type CreateResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type ErrResult struct {
	golangsdk.ErrResult
}

type ListRolesResult struct {
	golangsdk.Result
}

func (r ListRolesResult) ExtractRoles() ([]roles.Role, error) {
	var s struct {
		Roles []roles.Role `json:"roles"`
	}
	err := r.ExtractInto(&s)
	return s.Roles, err
}

type AgenciesPage struct {
	pagination.SinglePageBase
}

func ExtractAgencies(p pagination.Page) ([]Agency, error) {
	var agencies []Agency
	err := p.(AgenciesPage).ExtractIntoSlicePtr(&agencies, "agencies")
	if err != nil {
		return nil, err
	}
	return agencies, nil
}

func (p AgenciesPage) IsEmpty() (bool, error) {
	agencies, err := ExtractAgencies(p)
	return len(agencies) > 0, err
}
