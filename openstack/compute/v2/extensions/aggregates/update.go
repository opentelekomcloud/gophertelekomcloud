package aggregates

import (
	"strconv"

	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type UpdateOpts struct {
	// The name of the host aggregate.
	Name string `json:"name,omitempty"`
	// The availability zone of the host aggregate.
	// You should use a custom availability zone rather than
	// the default returned by the os-availability-zone API.
	// The availability zone must not include ‘:’ in its name.
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

// Update makes a request against the API to update a specific aggregate.
func Update(client *golangsdk.ServiceClient, aggregateID int, opts UpdateOpts) (*Aggregate, error) {
	b, err := build.RequestBody(opts, "aggregate")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL("os-aggregates", strconv.Itoa(aggregateID)), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
