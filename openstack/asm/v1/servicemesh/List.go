package servicemesh

import (
	golangsdk "github.com/opentelekomcloud/gophertelekomcloud"
	"github.com/opentelekomcloud/gophertelekomcloud/internal/extract"
)

// This function is used to obtain details about all service meshes.
func List(client *golangsdk.ServiceClient) ([]ServiceMesh, error) {
	// GET /v1/{project_id}/meshes
	raw, err := client.Get(client.ServiceURL("meshes"), nil, nil)
	if err != nil {
		return nil, err
	}

	var res ListMeshResponse
	err = extract.Into(raw.Body, &res)
	return res.Items, err
}

type ListMeshResponse struct {
	// API version. The value is fixed at v1 and cannot be changed.
	ApiVersion string `json:"apiVersion"`
	// API type. The value is fixed at MeshList and cannot be changed.
	Kind string `json:"kind"`
	// Service Mesh List
	Items []ServiceMesh `json:"items"`
}
