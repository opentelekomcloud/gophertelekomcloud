package projects

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	ID      string `q:"id"`
	Limit   int    `q:"limit"`
	Name    string `q:"name"`
	Offset  int    `q:"offset"`
	SortDir string `q:"sort_dir"`
	SortKey string `q:"sort_key"`
	Status  int    `q:"status"`
}

type ListResult struct {
	EnterpriseProjects []EnterpriseProject `json:"enterprise_projects"`
	TotalCount         int                 `json:"total_count"`
}

func List(client *golangsdk.ServiceClient, opts ListOpts) (*ListResult, error) {
	q, err := golangsdk.BuildQueryString(&opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("enterprise-projects")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResult
	err = extract.Into(raw.Body, &res)
	return &res, err
}
