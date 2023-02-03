package aggregates

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type CreateOpts struct {
	// The name of the host aggregate.
	Name string `json:"name" required:"true"`
	// The availability zone of the host aggregate.
	// You should use a custom availability zone rather than
	// the default returned by the os-availability-zone API.
	// The availability zone must not include ‘:’ in its name.
	AvailabilityZone string `json:"availability_zone,omitempty"`
}

// Create makes a request against the API to create an aggregate.
func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Aggregate, error) {
	b, err := build.RequestBody(opts, "aggregate")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("os-aggregates"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return extra(err, raw)
}
