package parameter

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ShowConfigurationDetail(client *golangsdk.ServiceClient, configId string) (*ShowConfigurationDetailResponse, error) {
	// GET https://{Endpoint}/v3/{project_id}/configurations/{config_id}
	raw, err := client.Get(client.ServiceURL("configurations", configId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ShowConfigurationDetailResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ShowConfigurationDetailResponse struct {
	// Parameter template ID
	Id string `json:"id,omitempty"`
	// Parameter template name
	Name string `json:"name,omitempty"`
	// Parameter template description
	Description string `json:"description,omitempty"`
	// DB version name
	DatastoreVersionName string `json:"datastore_version_name,omitempty"`
	// Database name
	DatastoreName string `json:"datastore_name,omitempty"`
	// Creation time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	Created string `json:"created,omitempty"`
	// Update time in the "yyyy-MM-ddTHH:mm:ssZ" format.
	Updated string `json:"updated,omitempty"`
	// The parameters defined by users based on the default parameter templates.
	ConfigurationParameters []ConfigurationParameterResult `json:"configuration_parameters,omitempty"`
}
