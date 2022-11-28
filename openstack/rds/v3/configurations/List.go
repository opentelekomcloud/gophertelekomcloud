package configurations

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List is used to obtain the parameter template list, including default
// parameter templates of all databases and those created by users.
func List(client *golangsdk.ServiceClient) ([]Configuration, error) {
	// GET https://{Endpoint}/v3/{project_id}/configurations
	raw, err := client.Get(client.ServiceURL("configurations"), nil, nil)

	var res []Configuration
	err = extract.IntoSlicePtr(raw.Body, &res, "configurations")
	return res, err
}
