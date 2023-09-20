package v3

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// ShowEngineVersion
// databaseName: DB engine. The following DB engine is supported (case-insensitive): gaussdb-mysql
func ShowEngineVersion(client *golangsdk.ServiceClient, databaseName string) ([]EngineVersionInfo, error) {
	// GET https://{Endpoint}/mysql/v3/{project_id}/datastores/{database_name}
	raw, err := client.Get(client.ServiceURL("datastores", databaseName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []EngineVersionInfo
	err = extract.IntoSlicePtr(raw.Body, &res, "datastores")
	return res, err
}

type EngineVersionInfo struct {
	// DB version ID. Its value is unique.
	Id string `json:"id"`
	// DB version number. Only the major version number with two digits is returned.
	Name string `json:"name"`
}
