package attachinterfaces

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// Get requests details on a single interface attachment by the server and port IDs.
func Get(client *golangsdk.ServiceClient, serverID, portID string) (*Interface, error) {
	raw, err := client.Get(client.ServiceURL("servers", serverID, "os-interface", portID), nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res Interface
	err = extract.IntoStructPtr(raw.Body, &res, "interfaceAttachment")
	return &res, err
}

// Interface represents a network interface on a server.
type Interface struct {
	PortState string    `json:"port_state"`
	FixedIPs  []FixedIP `json:"fixed_ips"`
	PortID    string    `json:"port_id"`
	NetID     string    `json:"net_id"`
	MACAddr   string    `json:"mac_addr"`
}
