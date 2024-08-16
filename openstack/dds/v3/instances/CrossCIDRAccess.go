package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type NetworkRangesOpts struct {
	// CIDR block where the client is located
	NetworkRanges []string `json:"client_network_ranges" required:"true"`
}

func CreateClientNetwork(client *golangsdk.ServiceClient, instanceId string, opts NetworkRangesOpts) (err error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/client-network
	_, err = client.Post(client.ServiceURL("instances", instanceId, "client-network"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
