package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Update makes a request against the API to update a specific aggregate.
func Update(client *golangsdk.ServiceClient, aggregateID int, opts UpdateOpts) (r UpdateResult) {
	v := strconv.Itoa(aggregateID)

	b, err := opts.ToAggregatesUpdateMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Put(aggregatesUpdateURL(client, v), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
