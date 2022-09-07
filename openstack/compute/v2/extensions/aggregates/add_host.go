package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// AddHost makes a request against the API to add host to a specific aggregate.
func AddHost(client *golangsdk.ServiceClient, aggregateID int, opts AddHostOpts) (r ActionResult) {
	v := strconv.Itoa(aggregateID)

	b, err := opts.ToAggregatesAddHostMap()
	if err != nil {
		return nil, err
	}
	raw, err := client.Post(client.ServiceURL("os-aggregates", v, "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
