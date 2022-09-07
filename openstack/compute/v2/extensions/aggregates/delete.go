package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
)

// Delete makes a request against the API to delete an aggregate.
func Delete(client *golangsdk.ServiceClient, aggregateID int) (err error) {
	_, err = client.Delete(client.ServiceURL("os-aggregates", strconv.Itoa(aggregateID)), &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
