package configurations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListConfigOpts struct {
	// Index offset.
	// If offset is set to N, the resource query starts from the N+1 piece of data. The default value is 0, indicating that the query starts from the first piece of data.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// Maximum number of specifications that can be queried
	// Value range: 1-100
	// If this parameter is not transferred, the first 100 pieces of specification information can be queried by default.
	Limit int `q:"limit"`
}

func List(client *golangsdk.ServiceClient, opts ListConfigOpts) (*ListResponse, error) {
	url, err := golangsdk.NewURLBuilder().WithEndpoints("configurations").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/configurations
	raw, err := client.Get(client.ServiceURL(url.String()), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListResponse struct {
	Configurations []ConfigurationsResponse `json:"configurations"`
	Quota          int                      `json:"quota"`
	TotalCount     int                      `json:"total_count"`
}

type ConfigurationsResponse struct {
	// Parameter template ID.
	ID string `json:"id"`
	// Parameter template name.
	Name string `json:"name"`
	// Parameter template description.
	Description string `json:"description"`
	// Database version.
	DatastoreVersion string `json:"datastore_version"`
	// Database type.
	DatastoreName string `json:"datastore_name"`
	// Node type of the parameter template.
	NodeType string `json:"node_type"`
	// Creation time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	CreatedAt string `json:"created"`
	// Update time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	UpdatedAt string `json:"updated"`
	// Indicates whether the parameter template is created by users.
	// false: The parameter template is a default parameter template.
	// true: The parameter template is a custom template.
	UserDefined bool `json:"user_defined"`
}
