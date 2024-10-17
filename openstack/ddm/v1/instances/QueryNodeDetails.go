package instances

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This API  is used to query details of a DDM instance node.
func QueryNodeDetails(client *golangsdk.ServiceClient, instanceId string, nodeId string) (*QueryNodeDetailsResponse, error) {

	// GET /v1/{project_id}/instances/{instance_id}/nodes/{node_id}
	raw, err := client.Get(client.ServiceURL("instances", instanceId, "nodes", nodeId), nil, nil)
	if err != nil {
		return nil, err
	}

	var res QueryNodeDetailsResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}

type QueryNodeDetailsResponse struct {
	// Node status. For details, see Status Description at:
	// https://docs.otc.t-systems.com/distributed-database-middleware/api-ref/appendix/status_description.html
	Status string `json:"status"`
	// Node name
	Name string `json:"name"`
	// Node ID
	NodeID string `json:"node_id"`
	// Private IP address of the node
	PrivateIP string `json:"private_ip"`
	// Floating IP address of the node
	FloatingIP string `json:"floating_ip"`
	// VM ID
	ServerID string `json:"server_id"`
	// Subnet name
	SubnetName string `json:"subnet_name"`
	// Data disk ID
	DataVolumeID string `json:"datavolume_id"`
	// IP address provided by the resource subnet
	ResSubnetIP string `json:"res_subnet_ip"`
	// System disk ID
	SystemVolumeID string `json:"systemvolume_id"`
}
