package parameter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListConfigurations(client *golangsdk.ServiceClient) (*ListConfigurationsResponse, error) {
	// GET https://{Endpoint}/v3/{project_id}/configurations
	raw, err := client.Get(client.ServiceURL("configurations"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListConfigurationsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListConfigurationsResponse struct {
	// Total number of queried records
	Count int32 `json:"count,omitempty"`
	// Parameter templates
	Configurations []ListConfigurationsResult `json:"configurations,omitempty"`
}

type ListConfigurationsResult struct {
	// Parameter template ID
	Id string `json:"id"`
	// Parameter template name
	Name string `json:"name"`
	// Parameter template description
	Description *string `json:"description,omitempty"`
	// DB version name
	DatastoreVersionName string `json:"datastore_version_name"`
	// Database name
	DatastoreName string `json:"datastore_name"`
	// Creation time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	Created string `json:"created"`
	// Update time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	Updated string `json:"updated"`
	// Whether the parameter template is a custom template.
	// false: The parameter template is a default parameter template.
	// true: The parameter template is a custom template.
	UserDefined bool `json:"user_defined"`
}
