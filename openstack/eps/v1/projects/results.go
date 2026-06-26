package projects

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type EnterpriseProject struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Status      int    `json:"status"`
	DeleteFlag  bool   `json:"delete_flag"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type commonResult struct {
	golangsdk.Result
}

func (r commonResult) Extract() (*EnterpriseProject, error) {
	var s struct {
		EnterpriseProject *EnterpriseProject `json:"enterprise_project"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return nil, err
	}
	return s.EnterpriseProject, nil
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

type ProjectPage struct {
	pagination.SinglePageBase
}

func (r ProjectPage) IsEmpty() (bool, error) {
	data, err := ExtractProjects(r)
	if err != nil {
		return false, err
	}
	return len(data.EnterpriseProjects) == 0, err
}

type ProjectListResult struct {
	EnterpriseProjects []EnterpriseProject `json:"enterprise_projects"`
	TotalCount         int                 `json:"total_count"`
}

func ExtractProjects(r pagination.Page) (*ProjectListResult, error) {
	var s ProjectListResult
	err := (r.(ProjectPage)).ExtractInto(&s)
	return &s, err
}
