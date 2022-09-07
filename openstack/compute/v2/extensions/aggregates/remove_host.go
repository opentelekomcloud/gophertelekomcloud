package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// RemoveHost makes a request against the API to remove host from a specific aggregate.
func RemoveHost(client *golangsdk.ServiceClient, aggregateID int, opts RemoveHostOpts) (r ActionResult) {
	v := strconv.Itoa(aggregateID)

	b, err := opts.ToAggregatesRemoveHostMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(aggregatesRemoveHostURL(client, v), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
