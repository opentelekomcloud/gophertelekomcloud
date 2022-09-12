package tenantnetworks

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get returns data about a previously created Network.
func Get(client *golangsdk.ServiceClient, id string) (*Network, error) {
	raw, err := client.Get(client.ServiceURL("os-tenant-networks", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var res Network
	err = extract.IntoStructPtr(raw.Body, &res, "network")
	return &res, err
}

// A Network represents a network that a server communicates on.
type Network struct {
	// CIDR is the IPv4 subnet.
	CIDR string `json:"cidr"`
	// ID is the UUID of the network.
	ID string `json:"id"`
	// Name is the common name that the network has.
	Name string `json:"label"`
}
