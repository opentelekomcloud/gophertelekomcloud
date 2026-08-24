package postgres_extensions

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListOpts struct {
	// Specifies the name of the specific database created inside the RDS instance.
	// This is the logical database name, not the RDS instance identifier.
	DatabaseName string `q:"database_name" required:"true"`
	// Specifies the index position. If offset is set to N, the resource query starts from the N+1 piece of data.
	// The value is 0 by default, indicating that the query starts from the first piece of data. The value must be a positive number.
	Offset int `q:"offset"`
	// Specifies the number of records to be queried. The default value is 100.
	// The value cannot be a negative number. The minimum value is 1 and the maximum value is 100.
	Limit int `q:"limit"`
}

// This function is used to list extensions for a specified database.
func List(client *golangsdk.ServiceClient, instanceId string, opts ListOpts) (*ListResponse, error) {
	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/extensions?database_name={database_name}&offset={offset}&limit={limit}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceId, "extensions").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Extensions []Extension `json:"extensions"`
	TotalCount int         `json:"total_count"`
}

type Extension struct {
	// Extension name.
	Name string `json:"name"`
	// The name of the specific database created inside the RDS instance.
	// This is the logical database name, not the RDS instance identifier.
	DatabaseName string `json:"database_name"`
	// Extension version.
	Version string `json:"version"`
	// New version that the extension can be upgraded to.
	// If the value of this parameter is different from that of Version,
	// the extension can be upgraded.
	VersionUpdate string `json:"version_update"`
	// Dependent preloaded library.
	SharedPreloadLibraries string `json:"shared_preload_libraries"`
	// Whether the extension has been created.
	Created bool `json:"created"`
	// Extension description.
	Description string `json:"description"`
	// Whether the extension can be installed.
	EnableInstall bool `json:"enable_install"`
}
