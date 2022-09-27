package v3

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListDatastores(client *golangsdk.ServiceClient, datastoreName string) ([]string, error) {
	// GET https://{Endpoint}/v3/{project_id}/datastores/{datastore_name}/versions
	raw, err := client.Get(client.ServiceURL("datastores", datastoreName, "versions"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []string
	err = extract.IntoSlicePtr(raw.Body, &res, "versions")
	return res, err
}
