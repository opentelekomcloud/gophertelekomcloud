package attachinterfaces

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts specifies parameters of a new interface attachment.
type CreateOpts struct {
	// PortID is the ID of the port for which you want to create an interface.
	// The NetworkID and PortID parameters are mutually exclusive.
	// If you do not specify the PortID parameter, the OpenStack Networking API
	// v2.0 allocates a port and creates an interface for it on the network.
	PortID string `json:"port_id,omitempty"`
	// NetworkID is the ID of the network for which you want to create an interface.
	// The NetworkID and PortID parameters are mutually exclusive.
	// If you do not specify the NetworkID parameter, the OpenStack Networking
	// API v2.0 uses the network information cache that is associated with the instance.
	NetworkID string `json:"net_id,omitempty"`
	// Slice of FixedIPs. If you request a specific FixedIP address without a
	// NetworkID, the request returns a Bad Request (400) response code.
	// Note: this uses the FixedIP struct, but only the IPAddress field can be used.
	FixedIPs []FixedIP `json:"fixed_ips,omitempty"`
}

// Create requests the creation of a new interface attachment on the server.
func Create(client *golangsdk.ServiceClient, serverID string, opts CreateOpts) (*Interface, error) {
	b, err := build.RequestBody(opts, "interfaceAttachment")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("servers", serverID, "os-interface"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Interface
	err = extract.IntoStructPtr(raw.Body, &res, "interfaceAttachment")
	return &res, err
}
