package instances

import (
	"bytes"

	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/pagination"
)

type QueryNodesOpts struct {
	// Specifies the Index offset.
	// The query starts from the next piece of data indexed by this parameter. The value is 0 by default.
	// The value must be a positive integer.
	Offset int `q:"offset"`
	// A maximum of nodes to be queried.
	// Value range: 1 to 128.
	// If the parameter value is not specified, 10 nodes are queried by default.
	Limit int `q:"limit"`
}

// QueryNodes function is used to query nodes of a DDM instance.
func QueryNodes(client *golangsdk.ServiceClient, instanceId string, opts QueryNodesOpts) ([]NodeList, error) {
	// GET /v1/{project_id}/instances/{instance_id}/nodes?offset={offset}&limit={limit}
	url, err := golangsdk.NewURLBuilder().WithEndpoints("instances", instanceId, "nodes").WithQueryParams(&opts).Build()
	if err != nil {
		return nil, err
	}

	pages, err := pagination.Pager{
		Client:     client,
		InitialURL: client.ServiceURL(url.String()),
		CreatePage: func(r pagination.NewPageResult) pagination.NewPage {
			return NodesPage{NewSinglePageBase: pagination.NewSinglePageBase{NewPageResult: r}}
		},
	}.NewAllPages()
	if err != nil {
		return nil, err
	}
	return ExtractNodes(pages)
}

func ExtractNodes(r pagination.NewPage) ([]NodeList, error) {
	var queryNodesResponse QueryNodesResponse
	err := extract.Into(bytes.NewReader((r.(NodesPage)).Body), &queryNodesResponse)
	return queryNodesResponse.Nodes, err
}

type NodesPage struct {
	pagination.NewSinglePageBase
}

type QueryNodesResponse struct {
	// Instance node information
	Nodes []NodeList `json:"nodes"`
	// Which page the server starts returning items.
	Offset int `json:"offset"`
	// Number of records displayed on each page
	Limit int `json:"limit"`
	// Number of DDM instance nodes
	Total int `json:"total"`
}

type NodeList struct {
	// Port
	Port string `json:"port"`
	// Node status.  For details, see Status Description at:
	// https://docs.otc.t-systems.com/distributed-database-middleware/api-ref/appendix/status_description.html
	Status string `json:"status"`
	// Node ID
	NodeID string `json:"node_id"`
	// IP address
	IP string `json:"ip"`
}
