package template

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func Get(client *golangsdk.ServiceClient, configId string) (*GetResponse, error) {
	raw, err := client.Get(client.ServiceURL("configurations", configId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res GetResponse
	return &res, extract.Into(raw.Body, &res)
}

type GetResponse struct {
	Id                      string                    `json:"id"`
	Name                    string                    `json:"name"`
	Description             string                    `json:"description"`
	DataStoreVersionName    string                    `json:"datastore_version_name"`
	DataStoreName           string                    `json:"datastore_name"`
	Created                 string                    `json:"created"`
	Updated                 string                    `json:"updated"`
	ConfigurationParameters []InstanceParameterResult `json:"configuration_parameters"`
}
