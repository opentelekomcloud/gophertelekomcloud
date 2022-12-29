package connection

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

type ListConnectionsOpts struct {
	NodeId string `q:"node_id"`
}

func ListConnections(client *golangsdk.ServiceClient, instanceId string, opts ListConnectionsOpts) (*ListConnectionsResponse, error) {
	q, err := golangsdk.BuildQueryString(opts)
	if err != nil {
		return nil, err
	}

	raw, err := client.Get(client.ServiceURL("instances", instanceId, "conn-statistics")+q.String(), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListConnectionsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type ListConnectionsResponse struct {
	TotalConnections      int           `json:"total_connections"`
	TotalInnerConnections int           `json:"total_inner_connections"`
	TotalOuterConnections int           `json:"total_outer_connections"`
	InnerConnections      []Connections `json:"inner_connections"`
	OuterConnections      []Connections `json:"outer_connections"`
}

type Connections struct {
	ClientIp string `json:"client_ip"`
	Count    int    `json:"count"`
}
