package attachinterfaces

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAttachInterfacesCreateMap() (map[string]interface{}, error)
}

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

// ToAttachInterfacesCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToAttachInterfacesCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "interfaceAttachment")
}
