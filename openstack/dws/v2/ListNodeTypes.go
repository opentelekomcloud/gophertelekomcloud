package v2

type ListNodeTypesRequest struct {
}

// GET /v2/{project_id}/node-types

type ListNodeTypesResponse struct {
	// List of node type objects
	NodeTypes []NodeTypes `json:"node_types,omitempty"`
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
