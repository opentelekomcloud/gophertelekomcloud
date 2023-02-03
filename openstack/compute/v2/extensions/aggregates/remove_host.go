package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type RemoveHostOpts struct {
	// The name of the host.
	Host string `json:"host" required:"true"`
}

// RemoveHost makes a request against the API to remove host from a specific aggregate.
func RemoveHost(client *golangsdk.ServiceClient, aggregateID int, opts RemoveHostOpts) (*Aggregate, error) {
	b, err := build.RequestBody(opts, "remove_host")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-aggregates", strconv.Itoa(aggregateID), "action"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
