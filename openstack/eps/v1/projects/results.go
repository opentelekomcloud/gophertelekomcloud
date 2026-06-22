package projects

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type EnterpriseProject struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Type        string `json:"type"`
	DeleteFlag  bool   `json:"delete_flag"`
}

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*EnterpriseProject, error) {
	var s struct {
		EnterpriseProject *EnterpriseProject `json:"enterprise_project"`
	}
	err := r.ExtractInto(&s)
	return s.EnterpriseProject, err
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	golangsdk.ErrResult
}

type ProjectPage struct {
	pagination.SinglePageBase
}

func (r ProjectPage) IsEmpty() (bool, error) {
	data, err := ExtractProjects(r)
	if err != nil {
		return false, err
	}
	return len(data) == 0, err
}

func ExtractProjects(r pagination.Page) ([]EnterpriseProject, error) {
	var s struct {
		EnterpriseProjects []EnterpriseProject `json:"enterprise_projects"`
	}
	err := (r.(ProjectPage)).ExtractInto(&s)
	return s.EnterpriseProjects, err
}
