package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get makes a request against the API to get details for a specific aggregate.
func Get(client *golangsdk.ServiceClient, aggregateID int) (*Aggregate, error) {
	raw, err := client.Get(client.ServiceURL("os-aggregates", strconv.Itoa(aggregateID)), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Aggregate
	err = extract.IntoStructPtr(raw.Body, &res, "aggregate")
	return &res, err
}
