package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get makes a request against the API to get details for a specific aggregate.
func Get(client *golangsdk.ServiceClient, aggregateID int) (*Aggregate, error) {
	raw, err := client.Get(client.ServiceURL("os-aggregates", strconv.Itoa(aggregateID)), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
