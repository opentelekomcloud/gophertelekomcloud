package instance

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
)

type ShrinkInstanceNodeOpts struct {
	InstanceId string
	// The number of nodes that are reduced randomly.
	// The value is 1.
	// NOTE
	// If the client is directly connected to a node, you are not advised to delete nodes randomly.
	Num int32 `json:"num,omitempty"`
	// The ID of the node to be deleted. The node must support scale-in. If this parameter is not transferred, the number of nodes to be deleted is specified based on the internal policy of the system.
	// NOTE
	// Either num or node_list must be set.
	// If num is transferred, the value must be 1.
	// If node_list is transferred, the value must be 1.
	// If the values of num and node_list are transferred at the same time, the value of node_list is used.
	// If the value of node_list is empty, scale-in is performed on random nodes.
	// If the value of node_list is not empty, the scale-in is performed based on the specified node ID.
	// Before scale-in, do not directly connect a node to prevent service interruption caused by scale-in.
	NodeList []string `json:"node_list,omitempty"`
}

func ShrinkInstanceNode(client *golangsdk.ServiceClient, opts ShrinkInstanceNodeOpts) (string, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return "", err
	}

	// POST https://{Endpoint}/v3/{project_id}/instances/{instance_id}/reduce-node
	raw, err := client.Post(client.ServiceURL("instances", opts.InstanceId, "reduce-node"), b, nil, nil)
	return extraJob(err, raw)
}
