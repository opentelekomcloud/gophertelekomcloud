package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func ListCollations(client *golangsdk.ServiceClient) ([]string, error) {
	// GET https://{Endpoint}/v3/{project_id}/collations
	raw, err := client.Get(client.ServiceURL("collations"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []string
	err = extract.IntoSlicePtr(raw.Body, &res, "charSets")
	return res, err
}
