package flavors

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

func List(client *golangsdk.ServiceClient) ([]Version, error) {
	raw, err := client.Get(client.ServiceURL("flavors"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Version
	err = extract.IntoSlicePtr(raw.Body, &res, "versions")
	return res, err
}
