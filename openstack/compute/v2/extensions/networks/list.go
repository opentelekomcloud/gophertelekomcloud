package networks

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List returns a Pager that allows you to iterate over a collection of Network.
func List(client *golangsdk.ServiceClient) ([]Network, error) {
	raw, err := client.Get(client.ServiceURL("os-networks"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Network
	err = extract.IntoSlicePtr(raw.Body, &res, "networks")
	return res, err
}
