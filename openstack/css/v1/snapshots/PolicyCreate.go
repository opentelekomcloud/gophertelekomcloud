package snapshots

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

// PolicyCreateOpts contains options for creating a snapshot policy.
// This object is passed to the snapshots.PolicyCreate function.
type PolicyCreateOpts struct {
	Prefix     string `json:"prefix" required:"true"`
	Period     string `json:"period" required:"true"`
	KeepDay    int    `json:"keepday" required:"true"`
	Enable     string `json:"enable" required:"true"`
	DeleteAuto string `json:"deleteAuto,omitempty"`
}

// PolicyCreate will create a new snapshot policy based on the values in PolicyCreateOpts.
func PolicyCreate(client *golangsdk.ServiceClient, opts PolicyCreateOpts, clusterId string) (err error) {
	b, err := build.RequestBodyMap(opts, "")
	if err != nil {
		return
	}

	_, err = client.Post(client.ServiceURL("clusters", clusterId, "index_snapshot/policy"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
