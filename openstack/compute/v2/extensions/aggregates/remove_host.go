package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// RemoveHost makes a request against the API to remove host from a specific aggregate.
func RemoveHost(client *golangsdk.ServiceClient, aggregateID int, opts RemoveHostOpts) (*Aggregate, error) {
	b, err := opts.ToAggregatesRemoveHostMap()
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-aggregates", strconv.Itoa(aggregateID), "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})

	if err != nil {
		return nil, err
	}

	var res Aggregate
	err = extract.IntoStructPtr(raw.Body, &res, "aggregate")
	return &res, err
}
