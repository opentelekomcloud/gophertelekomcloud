package configurations

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, configId string) (*Response, error) {
	// GET https://{Endpoint}/v3/{project_id}/configurations/{config_id}
	raw, err := client.Get(client.ServiceURL("configurations", configId), nil, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"},
	})
	if err != nil {
		return nil, err
	}

	var res Response
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type Response struct {
	// Parameter template ID.
	ID string `json:"id"`
	// Parameter template name.
	Name string `json:"name"`
	// Parameter template description.
	Description string `json:"description"`
	// Node type of the parameter template.
	NodeType string `json:"node_type"`
	// Database version.
	DatastoreVersion string `json:"datastore_version"`
	// Database type.
	DatastoreName string `json:"datastore_name"`
	// Creation time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	CreatedAt string `json:"created"`
	// Update time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	UpdatedAt string `json:"updated"`
	// The parameters defined by users based on the default parameter templates.
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	// The parameter name.
	Name string `json:"name"`
	// The parameter value.
	Value string `json:"value"`
	// The parameter description.
	Description string `json:"description"`
	// Parameter type.
	// The value can be integer, string, boolean, float, or list.
	Type string `json:"type"`
	// Value range.
	// For example, the value of integer is 0 or 1, and the value of boolean is true or false.
	ValueRange string `json:"value_range"`
	// Whether the instance needs to be restarted.
	// If the value is true, restart is required.
	// If the value is false, restart is not required.
	RestartRequired bool `json:"restart_required"`
	// Whether the parameter is read-only.
	// If the value is true, the parameter is read-only.
	// If the value false, the parameter is not read-only.
	ReadOnly bool `json:"readonly"`
}
