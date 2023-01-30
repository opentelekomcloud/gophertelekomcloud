package flavors

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListDatastores(client *golangsdk.ServiceClient, databaseName string) ([]DataStores, error) {
	// GET https://{Endpoint}/v3/{project_id}/datastores/{database_name}
	raw, err := client.Get(client.ServiceURL("datastores", databaseName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []DataStores
	err = extract.IntoSlicePtr(raw.Body, &res, "dataStores")
	return res, err
}

type DataStores struct {
	// Indicates the database version ID. Its value is unique.
	Id string `json:"id" `
	// Indicates the database version number. Only the major version number (two digits) is returned.
	// For example, if the version number is MySQL 5.6.X, only 5.6 is returned.
	Name string `json:"name"`
}
