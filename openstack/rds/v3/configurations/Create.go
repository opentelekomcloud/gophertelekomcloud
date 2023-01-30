package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// CreateOpts contains all the values needed to create a new configuration.
type CreateOpts struct {
	// Specifies the parameter template name. It contains a maximum of 64 characters and can contain only uppercase letters, lowercase letters, digits, hyphens (-), underscores (_), and periods (.).
	Name string `json:"name" required:"true"`
	// Specifies the parameter template description. It contains a maximum of 256 characters and cannot contain the following special characters: >!<"&'= Its value is left blank by default.
	Description string `json:"description,omitempty"`
	// Specifies the parameter values defined by users based on the default parameter template. By default, the parameter values cannot be changed.
	Values map[string]string `json:"values,omitempty"`
	// Specifies the database object.
	DataStore DataStore `json:"datastore" required:"true"`
}

type DataStore struct {
	// Specifies the DB engine. Its value can be any of the following and is case-insensitive:
	// MySQL
	// PostgreSQL
	// SQLServer
	Type string `json:"type" required:"true"`
	// Specifies the database version.
	// Example values:
	// MySQL: 8.0
	// PostgreSQL: 13
	// SQLServer: 2017_SE
	Version string `json:"version" required:"true"`
}

// Create will create a new Config based on the values in CreateOpts.
func Create(c *golangsdk.ServiceClient, opts CreateOpts) (*Configuration, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/configurations
	raw, err := c.Post(c.ServiceURL("configurations"), b, nil, &golangsdk.RequestOpts{
		OkCodes:     []int{200},
		MoreHeaders: openstack.StdRequestOpts().MoreHeaders,
	})
	if err != nil {
		return nil, err
	}

	var res Configuration
	err = extract.IntoStructPtr(raw.Body, &res, "configuration")
	return &res, err
}
