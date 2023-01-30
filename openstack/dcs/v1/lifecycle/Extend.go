package lifecycle

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// ExtendOpts is a struct which represents the parameters of extend function
type ExtendOpts struct {
	// New specification (memory space) of the DCS instance.
	// The new specification to which the DCS instance will be scaled up must be greater than the current specification.
	// Unit: GB.
	NewCapacity int `json:"new_capacity" required:"true"`
	// DCS instance specification code.
	// This parameter is optional for DCS Redis 3.0 instances.
	// This parameter is mandatory for DCS Redis 4.0 and Redis 5.0 instances.
	SpecCode string `json:"spec_code" required:"true"`
}

// Extend is extending for a dcs instance
func Extend(client *golangsdk.ServiceClient, id string, opts ExtendOpts) (err error) {
	body, err := build.RequestBody(opts, "")
	if err != nil {
		return
	}

	// POST /v1.0/{project_id}/instances/{instance_id}/extend
	_, err = client.Post(client.ServiceURL("instances", id, "extend"), body, nil, &golangsdk.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
