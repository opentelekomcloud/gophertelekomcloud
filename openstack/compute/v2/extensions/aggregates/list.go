package aggregates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// List makes a request against the API to list aggregates.
func List(client *golangsdk.ServiceClient) ([]Aggregate, error) {
	raw, err := client.Get(client.ServiceURL("os-aggregates"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res []Aggregate
	err = extract.IntoSlicePtr(raw.Body, &res, "aggregates")
	return res, err
}
