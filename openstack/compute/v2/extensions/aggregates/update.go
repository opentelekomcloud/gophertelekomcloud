package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Update makes a request against the API to update a specific aggregate.
func Update(client *golangsdk.ServiceClient, aggregateID int, opts UpdateOpts) (*Aggregate, error) {
	b, err := opts.ToAggregatesUpdateMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("os-aggregates", strconv.Itoa(aggregateID)), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Aggregate
	err = extract.IntoStructPtr(raw.Body, &res, "aggregate")
	return &res, err
}
