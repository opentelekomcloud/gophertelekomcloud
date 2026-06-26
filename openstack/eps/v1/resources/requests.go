package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type Resource struct {
	ProjectID    string `json:"project_id"`
	ProjectName  string `json:"project_name"`
	ResourceID   string `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

type ResourceError struct {
	ProjectID    string `json:"project_id"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
}

type FilterResult struct {
	Resources  []Resource      `json:"resources"`
	Errors     []ResourceError `json:"errors"`
	TotalCount int             `json:"total_count"`
}

type Match struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FilterOpts struct {
	Projects      []string `json:"projects,omitempty"`
	ResourceTypes []string `json:"resource_types"`
	Offset        int      `json:"offset,omitempty"`
	Limit         int      `json:"limit,omitempty"`
	Matches       []Match  `json:"matches,omitempty"`
}

// Filter queries resources belonging to an enterprise project.
func Filter(client *golangsdk.ServiceClient, projectID string, opts FilterOpts) (*FilterResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	url := client.ServiceURL("enterprise-projects", projectID, "resources", "filter")
	raw, err := client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res FilterResult
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type MigrateResource struct {
	ResourceID   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	RegionID     string `json:"region_id,omitempty"`
}

type MigrateOpts struct {
	ProjectID string            `json:"project_id"`
	Resources []MigrateResource `json:"resources"`
}

// Migrate moves resources from one enterprise project to another.
func Migrate(client *golangsdk.ServiceClient, projectID string, opts MigrateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	url := client.ServiceURL("enterprise-projects", projectID, "resources-migrate")
	_, err = client.Post(url, b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
