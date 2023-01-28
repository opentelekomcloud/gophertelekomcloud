package cluster

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack"
)

// ListNodeTypes Merge to common
func ListNodeTypes(client *golangsdk.ServiceClient) ([]NodeTypes, error) {
	// GET /v2/{project_id}/node-types
	raw, err := client.Get(client.ServiceURL("node-types"), nil, openstack.StdRequestOpts())
	if err != nil {
		return nil, err
	}

	var res []NodeTypes
	err = extract.IntoSlicePtr(raw.Body, &res, "node_types")
	return res, err
}

type NodeTypes struct {
	// Name of a node type
	SpecName string `json:"spec_name"`
	// Node type details
	Detail []Detail `json:"detail"`
	// Node type ID
	Id string `json:"id"`
}

type Detail struct {
	// Attribute type
	Type string `json:"type,omitempty"`
	// Attribute value
	Value string `json:"value"`
	// Attribute unit
	Unit string `json:"unit"`
}
