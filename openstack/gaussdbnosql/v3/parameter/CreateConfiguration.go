package parameter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateConfigurationOpts struct {
	// Parameter template name. It contains a maximum of 64 characters and can contain only uppercase letters,
	// lowercase letters, digits, hyphens (-), underscores (_), and periods (.).
	Name string `json:"name"`
	// Parameter template description. It contains a maximum of 256 characters except the following special characters:
	// >!<"&'= The value is left blank by default.
	Description string `json:"description,omitempty"`
	// Database object.
	Datastore ConfigurationDatastore `json:"datastore"`
	// The parameter values defined by users based on the default parameter template.
	// By default, the parameter values cannot be changed.
	Values map[string]string `json:"values,omitempty"`
}

type ConfigurationDatastore struct {
	// Database type. For GaussDB(for Cassandra), the value is "cassandra".
	Type string `json:"type"`
	// Database version The value 3.11 indicates that GaussDB(for Cassandra) 3.11 is supported.
	Version string `json:"version"`
}

func CreateInstance(client *golangsdk.ServiceClient, opts CreateConfigurationOpts) (*CreateConfigurationResult, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST https://{Endpoint}/v3/{project_id}/configurations
	raw, err := client.Post(client.ServiceURL("configurations"), b, nil, nil)
	if err != nil {
		return nil, err
	}

	var res CreateConfigurationResult
	err = extract.IntoStructPtr(raw.Body, &res, "configuration")
	return &res, err
}

type CreateConfigurationResult struct {
	// Parameter template ID
	Id string `json:"id"`
	// Parameter template name
	Name string `json:"name"`
	// Parameter template description.
	Description string `json:"description,omitempty"`
	// DB version name
	DatastoreVersionName string `json:"datastore_version_name"`
	// Database name
	DatastoreName string `json:"datastore_name"`
	// Creation time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	Created string `json:"created"`
	// Update time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	Updated string `json:"updated"`
}
