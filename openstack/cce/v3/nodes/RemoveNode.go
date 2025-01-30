package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// RemoveNodesOpts contains parameters for removing nodes from a cluster
type RemoveNodesOpts struct {
	ClusterID string `json:"-"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion,omitempty"`
	// API type, fixed value RemoveNodesTask
	Kind string `json:"kind,omitempty"`
	// Configuration information
	Spec RemoveNodesSpec `json:"spec" required:"true"`
}

// RemoveNodesSpec contains the configuration for node removal
type RemoveNodesSpec struct {
	// Node login mode
	Login LoginSpec `json:"login" required:"true"`
	// List of nodes to be operated
	Nodes []NodeItem `json:"nodes" required:"true"`
}

// NodeItem represents a node to be removed
type NodeItem struct {
	// Node ID
	UID string `json:"uid" required:"true"`
}

// RemoveNodesResponse represents the response from the remove nodes operation
type RemoveNodesResponse struct {
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion,omitempty"`
	// API type, fixed value RemoveNodesTask
	Kind string `json:"kind,omitempty"`
	// Configuration information
	Spec RemoveNodesSpec `json:"spec"`
	// Job status
	Status TaskStatus `json:"status"`
}

// TaskStatus contains the job status information
type TaskStatus struct {
	// Job ID for tracking the removal process
	JobID string `json:"jobID"`
}

// Remove sends a request to remove nodes from the specified cluster
func Remove(client *golangsdk.ServiceClient, opts RemoveNodesOpts) (*RemoveNodesResponse, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// PUT /api/v3/projects/{project_id}/clusters/{cluster_id}/nodes/operation/remove
	raw, err := client.Put(client.ServiceURL("clusters", opts.ClusterID, "nodes", "operation", "remove"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	if err != nil {
		return nil, err
	}

	var res RemoveNodesResponse
	err = extract.Into(raw.Body, &res)
	return &res, err
}
