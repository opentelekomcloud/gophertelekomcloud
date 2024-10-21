package connection_monitoring

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type CreateOpts struct {
	// Specifies the ID of the VPN connection to be monitored.
	ConnectionID string `json:"vpn_connection_id" required:"true"`
}

func Create(client *golangsdk.ServiceClient, opts CreateOpts) (*Monitor, error) {
	b, err := build.RequestBody(opts, "connection_monitor")
	if err != nil {
		return nil, err
	}

	raw, err := client.Post(client.ServiceURL("connection-monitors"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202, 201},
	})
	if err != nil {
		return nil, err
	}

	var res Monitor
	return &res, extract.IntoStructPtr(raw.Body, &res, "connection_monitor")
}

type Monitor struct {
	// Specifies the ID of a VPN connection monitor.
	ID string `json:"id"`
	// Specifies the ID of the VPN connection to be monitored.
	ConnectionId string `json:"vpn_connection_id"`
	// Specifies the type of objects to be monitored.
	Type string `json:"type"`
	// Specifies the source address to be monitored.
	SourceIp string `json:"source_ip"`
	// Specifies the destination address to be monitored.
	DestinationIp string `json:"destination_ip"`
	// Specifies the protocol used by NQA.
	ProtoType string `json:"proto_type"`
	// Specifies the status of the VPN connection monitor.
	// Value range:
	// ACTIVE: normal
	// PENDING_CREATE: creating
	// PENDING_DELETE: deleting
	Status string `json:"status"`
}
