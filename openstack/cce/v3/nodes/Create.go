package nodes

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/build"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// CreateOpts is a struct contains the parameters of creating Node
type CreateOpts struct {
	// API type, fixed value Node
	Kind string `json:"kind" required:"true"`
	// API version, fixed value v3
	ApiVersion string `json:"apiVersion" required:"true"`
	// Metadata required to create a Node
	Metadata CreateMetaData `json:"metadata"`
	// specifications to create a Node
	Spec Spec `json:"spec" required:"true"`
}

// CreateMetaData required to create a Node
type CreateMetaData struct {
	// Node name
	Name string `json:"name,omitempty"`
	// Node tag, key value pair format
	Labels map[string]string `json:"labels,omitempty"`
	// Node annotation, key value pair format
	Annotations map[string]string `json:"annotations,omitempty"`
}

// Create accepts a CreateOpts struct and uses the values to create a new
// logical node.
func Create(client *golangsdk.ServiceClient, clusterID string, opts CreateOpts) (*Nodes, error) {
	b, err := build.RequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	// POST /api/v3/projects/{project_id}/clusters/{cluster_id}/nodes
	raw, err := client.Post(client.ServiceURL("clusters", clusterID, "nodes"), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{201},
	})
	if err != nil {
		return nil, err
	}

	var res Nodes
	return &res, extract.Into(raw.Body, &res)
}
