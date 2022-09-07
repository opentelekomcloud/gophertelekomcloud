package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete makes a request against the API to delete an aggregate.
func Delete(client *golangsdk.ServiceClient, aggregateID int) (r DeleteResult) {
	v := strconv.Itoa(aggregateID)
	raw, err := client.Delete(client.ServiceURL("os-aggregates", v), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
