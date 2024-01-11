package v3

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListConfigOpts struct {
	// Index offset. If offset is set to N, the resource query starts from the N+1 piece of data.
	// The default value is 0, indicating that the query starts from the first piece of data.
	// The value must be a positive integer
	Offset int `q:"offset"`
	// Number of records to be queried. The default value is 100.
	// The value must be a positive integer.
	// The minimum value is 1 and the maximum value is 100.
	Limit int `q:"limit"`
}

func ListConfigurations(client *golangsdk.ServiceClient, opts ListConfigOpts) (*ListConfigResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/mysql/v3/{project_id}/configurations
	raw, err := client.Get(client.ServiceURL("configurations")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListConfigResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListConfigResponse struct {
	// Parameter template information
	Configurations []ConfigurationSummary `json:"configurations"`
	// Total number of parameter templates
	TotalCount int32 `json:"total_count"`
}

type ConfigurationSummary struct {
	// Parameter template ID
	Id string `json:"id"`
	// Parameter template name
	Name string `json:"name"`
	// Parameter template description
	Description string `json:"description"`
	// DB version name
	DatastoreVersionName string `json:"datastore_version_name"`
	// Database name
	DatastoreName string `json:"datastore_name"`
	// Creation time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time.
	// Z indicates the time zone offset.
	// For example, for French Winter Time (FWT), the time offset is shown as +0200.
	Created string `json:"created"`
	// Update time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	// T is the separator between the calendar and the hourly notation of time.
	// Z indicates the time zone offset.
	// For example, for French Winter Time (FWT), the time offset is shown as +0200.
	Updated string `json:"updated"`
	// Whether the parameter template is a custom template.
	// false: The parameter template is a default template.
	// true: The parameter template is a custom template.
	UserDefined bool `json:"user_defined"`
}
