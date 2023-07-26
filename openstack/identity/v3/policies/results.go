package policies

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type commonResult struct {
	golangsdk.Result
}

type ListResult struct {
	commonResult
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

type DeleteResult struct {
	golangsdk.ErrResult
}

type ListPolicy struct {
	Links       Links    `json:"links"`
	Roles       []Policy `json:"roles"`
	TotalNumber int      `json:"total_number"`
}

type Links struct {
	Self     string `json:"self"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

type Policy struct {
	DomainId    string       `json:"domain_id"`
	References  int          `json:"references,omitempty"`
	UpdatedTime string       `json:"updated_time,omitempty"`
	CreatedTime string       `json:"created_time,omitempty"`
	Catalog     string       `json:"catalog"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Links       RolesLinks   `json:"links"`
	ID          string       `json:"id"`
	DisplayName string       `json:"display_name"`
	Type        string       `json:"type"`
	Policy      CustomPolicy `json:"policy"`
}

type RolesLinks struct {
	Self string `json:"self"`
}

type PolicyPage struct {
	pagination.LinkedPageBase
}

type CustomPolicy struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Action    []string  `json:"Action"`
	Effect    string    `json:"Effect"`
	Condition Condition `json:"Condition,omitempty"`
	Resource  any       `json:"Resource,omitempty"`
}

func (r commonResult) ExtractPolicies() (ListPolicy, error) {
	var s ListPolicy
	err := r.ExtractIntoStructPtr(&s, "")
	if err != nil {
		return s, err
	}

	return s, nil
}

func (r commonResult) Extract() (*Policy, error) {
	var s struct {
		Policy *Policy `json:"role"`
	}
	err := r.ExtractInto(&s)
	return s.Policy, err
}
