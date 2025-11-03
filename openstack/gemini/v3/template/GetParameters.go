package template

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func GetInstanceParameters(client *golangsdk.ServiceClient, instanceId string) (*ParameterResponse, error) {
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "configurations"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ParameterResponse
	return &res, extract.Into(raw.Body, &res)
}

type ParameterResponse struct {
	DataStoreVersionName    string                    `json:"datastore_version_name"`
	DataStoreName           string                    `json:"datastore_name"`
	Created                 string                    `json:"created"`
	Updated                 string                    `json:"updated"`
	Id                      string                    `json:"id"`
	Mode                    string                    `json:"mode"`
	ConfigurationParameters []InstanceParameterResult `json:"configuration_parameters"`
}

type InstanceParameterResult struct {
	Name            string `json:"name"`
	Value           string `json:"value"`
	RestartRequired bool   `json:"restart_required"`
	Readonly        bool   `json:"readonly"`
	ValueRange      string `json:"value_range"`
	Type            string `json:"type"`
	Description     string `json:"description"`
}
