package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"net/url"
)

type ListConnectionsOpts struct {
	// Specifies the DB instance ID.
	InstanceId string `json:"-"`
	// Specifies the node ID.
	// If this parameter is left blank, the number of connections of all nodes that can be connected in the instance is queried.
	NodeId string `q:"node_id"`
}

func ListConnections(client *golangsdk.ServiceClient, opts ListConnectionsOpts) (*ListConnectionsResponse, error) {
	var opts2 interface{} = opts
	q, err := build.QueryString(opts2)
	if err != nil {
		return nil, err
	}

	// GET https://{Endpoint}/v3/{project_id}/instances/{instance_id}/conn-statistics
	raw, err := client.Get(client.ServiceURL("instances", opts.InstanceId, "conn-statistics")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListConnectionsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListConnectionsResponse struct {
	// Indicates the total number of connections, including internal and external connections.
	TotalConnections int `json:"total_connections"`
	// Indicates the total number of internal connections.
	TotalInnerConnections int `json:"total_inner_connections"`
	// Indicates the total number of external connections.
	TotalOuterConnections int `json:"total_outer_connections"`
	// Indicates the internal connection statistics array. Up to 200 records are supported.
	InnerConnections []Connections `json:"inner_connections"`
	// Indicates the external connection statistics array. Up to 200 records are supported.
	OuterConnections []Connections `json:"outer_connections"`
}

type Connections struct {
	// Indicates the IP address of the client connected to the instance or node.
	ClientIp string `json:"client_ip"`
	// Indicates the number of connections corresponding to the IP address.
	Count int `json:"count"`
}
