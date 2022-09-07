package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Get makes a request against the API to get details for a specific aggregate.
func Get(client *golangsdk.ServiceClient, aggregateID int) (r GetResult) {
	v := strconv.Itoa(aggregateID)
	raw, err := client.Get(aggregatesGetURL(client, v), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
