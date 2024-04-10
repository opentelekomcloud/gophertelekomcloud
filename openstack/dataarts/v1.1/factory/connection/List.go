package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

const (
	URLParamOffset         = "offset"
	URLParamLimit          = "limit"
	URLParamConnectionName = "connectionName"
)

// List is used to query a connection list.
// Send request GET /v1/{project_id}/connections?offset={offset}&limit={limit}&connectionName={connectionName}
func List(client *golangsdk.ServiceClient, urlParams map[string]string) ([]*Config, error) {

	raw, err := client.Get(client.ServiceURL("clusters"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []*Config
	err = extract.IntoSlicePtr(raw.Body, &res, "clusters")
	return res, err
}
