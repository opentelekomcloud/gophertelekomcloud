package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type MigrateResource struct {
	ResourceID   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	RegionID     string `json:"region_id,omitempty"`
}

type MigrateOpts struct {
	ProjectID string            `json:"project_id"`
	Resources []MigrateResource `json:"resources"`
}

func Migrate(client *golangsdk.ServiceClient, projectID string, opts MigrateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("enterprise-projects", projectID, "resources-migrate"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
