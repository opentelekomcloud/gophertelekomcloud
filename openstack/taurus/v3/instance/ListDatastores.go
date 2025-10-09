package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListDatastores(client *golangsdk.ServiceClient, databaseName string) (*DatastoresResponse, error) {
	raw, err := client.Get(client.ServiceURL("datastores", databaseName), nil, nil)
	if err != nil {
		return nil, err
	}

	var res DatastoresResponse
	err = extract.IntoStructPtr(raw.Body, &res, "")
	return &res, err
}

type DatastoresResponse struct {
	Datastores []MysqlEngineVersionInfo `json:"datastores"`
}

type MysqlEngineVersionInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
