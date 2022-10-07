package v2

type ListNodeTypesRequest struct {
}

// GET /v2/{project_id}/node-types

type ListNodeTypesResponse struct {
	//
	NodeTypes []NodeTypes `json:"node_types,omitempty"`
}

type NodeTypes struct {
	//
	SpecName string `json:"spec_name"`
	//
	Detail []Detail `json:"detail"`
	//
	Id string `json:"id"`
}

type Detail struct {
	//
	Type string `json:"type,omitempty"`
	//
	Value string `json:"value"`
	//
	Unit string `json:"unit"`
}
