package resources

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type MigrateOpts struct {
	ProjectID    string `json:"project_id,omitempty"`
	ResourceID   string `json:"resource_id" required:"true"`
	ResourceType string `json:"resource_type" required:"true"`
	RegionID     string `json:"region_id,omitempty"`
	Associated   bool   `json:"associated,omitempty"`
}

// Migrate moves a single resource to the target enterprise project.
func Migrate(client *golangsdk.ServiceClient, enterpriseProjectID string, opts MigrateOpts) error {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	_, err = client.Post(client.ServiceURL("enterprise-projects", enterpriseProjectID, "resources-migrate"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return err
}
