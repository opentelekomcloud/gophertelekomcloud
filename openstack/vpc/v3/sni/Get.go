package sni

import (
	"github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// Get is used to query details about a supplementary network interface.
func Get(client *golangsdk.ServiceClient, id string) (*SubNetworkInterface, error) {
	// GET /v3/{project_id}/vpc/sub-network-interfaces/{sub_network_interface_id}
	raw, err := client.Get(client.ServiceURL("sub-network-interfaces", id), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res SubNetworkInterface
	err = extract.IntoStructPtr(raw.Body, &res, "sub_network_interface")
	return &res, err
}

type SubNetworkInterfaceCount struct {
	// Request ID
	RequestId string `json:"request_id"`
	// Number of supplementary network interfaces
	SubNetworkInterfaces int `json:"sub_network_interfaces"`
}

// GetCount is used to query the number of supplementary network interfaces.
func GetCount(client *golangsdk.ServiceClient) (*SubNetworkInterfaceCount, error) {
	// GET /v3/{project_id}/vpc/sub-network-interfaces/count
	raw, err := client.Get(client.ServiceURL("sub-network-interfaces", "count"), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res SubNetworkInterfaceCount
	err = extract.Into(raw.Body, &res)
	return &res, err
}
