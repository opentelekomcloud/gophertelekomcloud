package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type UpdateOpts struct {
	// Central network ID.
	CentralNetworkId string `json:"-" required:"true"`
	// ID of the connection on the central network.
	ConnectionId string `json:"-" required:"true"`
	// Bandwidth type. Value options: BandwidthPackage, TestBandwidth.
	BandwidthType string `json:"bandwidth_type" required:"true"`
	// Bandwidth size in Mbit/s. Mandatory if bandwidth_type is BandwidthPackage.
	BandwidthSize int64 `json:"bandwidth_size,omitempty"`
	// Global connection bandwidth ID. Mandatory if bandwidth_type is BandwidthPackage.
	GlobalConnectionBandwidthId string `json:"global_connection_bandwidth_id,omitempty"`
}

func Update(client *golangsdk.ServiceClient, opts UpdateOpts) (*CentralNetworkConnection, error) {
	b, err := build.RequestBody(opts, "central_network_connection")
	if err != nil {
		return nil, err
	}

	raw, err := client.Put(client.ServiceURL(client.DomainID, "gcn", "central-network", opts.CentralNetworkId, "connections", opts.ConnectionId), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200, 202},
	})
	if err != nil {
		return nil, err
	}

	var res struct {
		CentralNetworkConnection CentralNetworkConnection `json:"central_network_connection"`
	}
	err = extract.Into(raw.Body, &res)
	return &res.CentralNetworkConnection, err
}
