package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List is used to obtain the parameter template list, including default
// parameter templates of all databases and those created by users.
func List(client *golangsdk.ServiceClient) ([]Configuration, error) {
	// GET https://{Endpoint}/v3/{project_id}/configurations
	raw, err := client.Get(client.ServiceURL("configurations"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Configuration
	err = extract.IntoSlicePtr(raw.Body, &res, "configurations")
	return res, err
}

type Configuration struct {
	// Specifies the parameter template ID.
	ID string `json:"id"`
	// Indicates the parameter template name.
	Name string `json:"name"`
	// Indicates the database version name.
	DatastoreVersionName string `json:"datastore_version_name"`
	// Indicates the database name.
	DatastoreName string `json:"datastore_name"`
	// Indicates the parameter template description.
	Description string `json:"description"`
	// Indicates the creation time in the following format: yyyy-MM-ddTHH:mm:ssZ.
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	Created string `json:"created"`
	// Indicates the update time in the following format: yyyy-MM-ddTHH:mm:ssZ.
	// T is the separator between the calendar and the hourly notation of time. Z indicates the time zone offset.
	Updated string `json:"updated"`
	// Specifies whether the parameter template is created by users.
	// false: The parameter template is a default parameter template.
	// true: The parameter template is a custom template.
	UserDefined bool `json:"user_defined"`
	// Configuration Parameters
	Parameters []Parameter `json:"configuration_parameters"`
}
